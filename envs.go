package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Env struct {
	Name string
	Shell
	OS
}

func RetrieveEnv(name string) *Env {
	env := &Env{Name: name}

	env.Retrieve()

	return env
}

func (e Env) Home() string {
	return os.Getenv("HOME")
	//return e.OS.GetUsersDir() + "/" + e.Name
}

func (e *Env) Retrieve() {
	e.OS = GetOS()

	e.OS.CheckUser(e.Name)

	e.Shell = GetShell(e.OS.GetShellPath(e.Name))
}

func (env *Env) LoadToVNEConfig(path string) {
	//config := CreateConfig(path)

	updateVNEConfig(path, env)

	log.Print("updated config file: " + path)
}

func (e *Env) UnloadToUser(path string) {
	userConf := CreateEmptyConfig()

	userConf.SetEnvVars(GetUserEnvVars(e))
	userConf.SetTools(GetUserTools(e))

	fileConig := CreateConfig(path)

	diff := fileConig.Minus(userConf)

	UnloadEnvVarsToUser(e, diff)
	UploadLackingTools(e, diff)
	CreateDirs(e, diff)
}

func CreateDirs(e *Env, conf *Config) {
	for _, dir := range conf.GetDirs() {
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

func updateVNEConfig(configPath string, env *Env) {
	envVars := GetUserEnvVars(env)
	tools := GetUserTools(env)

	content := "[envs]\n"

	for key, value := range envVars {
		content += key + "=" + value + "\n"
	}

	content += "[tools]\n"

	for _, tool := range tools {
		content += tool + "\n"
	}

	err := os.WriteFile(configPath, []byte(content), 0644) // TODO: deal with permissions
	Check(err)
}
