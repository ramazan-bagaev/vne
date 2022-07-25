package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type Env struct {
	Name         string
	EnvVariables map[string]string
}

func CreateEnv(name string) *Env {
	return &Env{Name: name}
}

func (e Env) ConfigPath() string {
	return "/Users/" + e.Name + "/.vne"
}

func (e Env) Home() string {
	return "/Users/" + e.Name
}

func (e *Env) Retrieve() {
	_, err := os.Stat(e.Home())

	if err != nil {
		CreateOSUser(e.Name)
	}

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

	configEnvs := make(map[string]string)
	for _, line := range configs["envs"] {
		split := strings.Split(line, "=")
		if len(split) != 2 {
			log.Println("can't parse env variable")
			return
		}

		configEnvs[split[0]] = split[1]
	}

	e.EnvVariables = configEnvs
}

func (e *Env) LoadToConfig() {
	LoadEnvVarsFromUser(e)
}

func (e Env) PrintEnvs() {
	for k, v := range e.EnvVariables {
		fmt.Println(k + "=" + v)
	}
}
