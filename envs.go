package main

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Env struct {
	Name       string
	ConfigPath string
	EnvVars    map[string]string
	Tools      []string
	Dirs       []string
	Shell
	OS
}

func CreateEnv(name string, config string) *Env {
	env := &Env{Name: name, ConfigPath: config}

	env.Retrieve()

	return env
}

func (e Env) Home() string {
	return e.OS.GetUsersDir() + "/" + e.Name
}

func (e *Env) Retrieve() {
	e.OS = GetOS()
	_, err := os.Stat(e.Home()) // TODO: check real users, not directories in /home or whatever

	if err != nil {
		log.Fatalf("should create env first: 'sudo vne -create %s', then login as a new user: 'login %s'", e.Name, e.Name)
	}

	e.OS.CheckUser(e.Name)

	_, err = os.Stat(e.ConfigPath)

	if err != nil {
		_, err = os.Create(e.ConfigPath)

		Check(err)
	}

	configLines := ParseFile(e.ConfigPath)
	configs := make(map[string][]string)
	currentConfName := "undefined"
	for _, line := range configLines {
		if match, _ := regexp.MatchString("\\[.", line); match {
			currentConfName = line[1 : len(line)-1]
		} else {
			configs[currentConfName] = append(configs[currentConfName], line)
		}
	}

	e.EnvVars = GetEnvVars(configs)
	e.Tools = GetTools(configs)
	e.Dirs = configs["dirs"]
	e.Shell = GetShell(e.OS.GetShellPath(e.Name))
}

func (e *Env) LoadToVNEConfig() {
	e.EnvVars = GetUserEnvVars(e)
	e.Tools = GetUserTools(e)

	e.PrintEnvs()
	e.PrintTools()
	e.PrintDirs()

	updateVNEConfig(e)
}

func (e *Env) UnloadToUser() {
	userEnv := CreateEnv(e.Name, e.ConfigPath)
	userEnv.EnvVars = GetUserEnvVars(userEnv)
	userEnv.Tools = GetUserTools(userEnv)

	d := e.Remove(userEnv)

	d.PrintEnvs()
	d.PrintTools()

	UnloadEnvVarsToUser(d)
	UploadLackingTools(d)
	CreateDirs(e)
}

func CreateDirs(e *Env) {
	for _, dir := range e.Dirs {
		if _, err := os.Stat(dir); err != nil {
			home := e.Home()
			var path string
			if dir == "~" {
				path = home
			} else if strings.HasPrefix(dir, "~/") {
				path = filepath.Join(home, dir[2:])
			}
			log.Print("creating dir: " + path)
			err := os.Mkdir(path, 0744) // TODO: have no idea which permissions
			Check(err)
		}
	}
}

func (e1 *Env) Remove(e2 *Env) *Env {
	if e1.Name != e2.Name {
		panic("remove method on env work only for same name environment")
	}

	res := CreateEnv(e1.Name, e1.ConfigPath)

	resVars := make(map[string]string)
	for e1K, e1V := range e1.EnvVars {
		if _, ok := e2.EnvVars[e1K]; !ok {
			resVars[e1K] = e1V
		}
	}

	resTools := []string{}
	for _, e1T := range e1.Tools {
		uniq := true
		for _, e2T := range e2.Tools {
			if e1T == e2T {
				uniq = false
			}
		}

		if uniq {
			resTools = append(resTools, e1T)
		}
	}

	res.EnvVars = resVars
	res.Tools = resTools

	return res
}

func (e Env) PrintEnvs() {
	log.Print("printing all env variables:")
	for k, v := range e.EnvVars {
		log.Print(k + "=" + v)
	}
}

func (e Env) PrintTools() {
	log.Print("printing all tools:")
	for _, el := range e.Tools {
		log.Print(el)
	}
}

func (e Env) PrintDirs() {
	log.Print("printing all dirs:")
	for _, el := range e.Dirs {
		log.Print(el)
	}
}

func updateVNEConfig(env *Env) {
	content := "[envs]\n"

	for key, value := range env.EnvVars {
		content += key + "=" + value + "\n"
	}

	content += "[tools]\n"

	for _, tool := range env.Tools {
		content += tool + "\n"
	}

	content += "[dirs]\n"

	for _, dir := range env.Dirs {
		content += dir + "\n"
	}

	err := os.WriteFile(env.ConfigPath, []byte(content), 0644) // TODO: deal with permissions
	Check(err)
}
