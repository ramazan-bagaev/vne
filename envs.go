package main

import (
	"log"
	"os"
	"regexp"
)

type Env struct {
	Name         string
	EnvVariables map[string]string
	Tools        []string
	Shell
	OS
}

func CreateEnv(name string) *Env {
	env := &Env{Name: name}

	env.Retrieve()

	return env
}

func (e Env) ConfigPath() string {
	return e.Home() + "/.vne"
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

	_, err = os.Stat(e.ConfigPath())

	if err != nil {
		_, err = os.Create(e.ConfigPath())

		Check(err)
	}

	configLines := ParseFile(e.ConfigPath())
	configs := make(map[string][]string)
	currentConfName := "undefined"
	for _, line := range configLines {
		if match, _ := regexp.MatchString("\\[.", line); match {
			currentConfName = line[1 : len(line)-1]
		} else {
			configs[currentConfName] = append(configs[currentConfName], line)
		}
	}

	e.EnvVariables = GetEnvVars(configs)
	e.Tools = GetTools(configs)
	e.Shell = GetShell(e.OS.GetShellPath(e.Name))
}

func (e *Env) LoadToVNEConfig() {
	e.EnvVariables = GetUserEnvVars(e)
	e.Tools = GetUserTools(e)

	e.PrintEnvs()
	e.PrintTools()

	updateVNEConfig(e)
}

func (e *Env) UnloadToUser() {
	userEnv := CreateEnv(e.Name)
	userEnv.EnvVariables = GetUserEnvVars(userEnv)
	userEnv.Tools = GetUserTools(userEnv)

	d := e.Remove(userEnv)

	d.PrintEnvs()
	d.PrintTools()

	UnloadEnvVarsToUser(d)
	UploadLackingTools(d)
}

func (e1 *Env) Remove(e2 *Env) *Env {
	if e1.Name != e2.Name {
		panic("remove method on env work only for same name environment")
	}

	res := CreateEnv(e1.Name)

	resVars := make(map[string]string)
	for e1K, e1V := range e1.EnvVariables {
		if _, ok := e2.EnvVariables[e1K]; !ok {
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

	res.EnvVariables = resVars
	res.Tools = resTools

	return res
}

func (e Env) PrintEnvs() {
	log.Print("printing all env variables:")
	for k, v := range e.EnvVariables {
		log.Print(k + "=" + v)
	}
}

func (e Env) PrintTools() {
	log.Print("printing all tools:")
	for _, el := range e.Tools {
		log.Print(el)
	}
}

func updateVNEConfig(env *Env) {
	content := "[envs]\n"

	for key, value := range env.EnvVariables {
		content += key + "=" + value + "\n"
	}

	content += "[tools]\n"

	for _, tool := range env.Tools {
		content += tool + "\n"
	}

	err := os.WriteFile(env.ConfigPath(), []byte(content), 0644)
	Check(err)
}
