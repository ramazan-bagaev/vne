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
	Console
	OS
}

func CreateEnv(name string) *Env {
	env := &Env{Name: name}

	env.Retrieve()

	return env
}

func (e Env) ConfigPath() string {
	return "/Users/" + e.Name + "/.vne"
}

func (e Env) Home() string {
	return "/Users/" + e.Name
}

func (e *Env) Retrieve() {
	_, err := os.Stat(e.Home()) // TODO: check real users, not directories in /home or whatever

	if err != nil {
		log.Fatalf("should create env first: 'sudo vne -create %s', then login as a new user: 'login %s'", e.Name, e.Name)
	}

	e.OS = GetOS()

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
	e.Console = GetConsole(e.OS.GetShellPath(e.Name))
}

func (e *Env) LoadToVNEConfig() {
	e.EnvVariables = GetUserEnvVars(e)
	e.Tools = GetUserTools(e)

	updateVNEConfig(e)
}

func (e *Env) UnloadToUser() {
	UnloadEnvVarsToUser(e)
	UploadLackingTools(e)
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
