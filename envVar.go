package main

import (
	"log"
	"os"
	"strings"
)

func GetEnvVars(configs map[string][]string) map[string]string {
	envVars := make(map[string]string)
	for _, line := range configs["envs"] {
		split := strings.Split(line, "=")
		if len(split) != 2 {
			log.Println("can't parse env variable")
			panic(1)
		}

		envVars[split[0]] = split[1]
	}

	return envVars
}

func LoadEnvVarsToVNE(env *Env) {
	envVars := make(map[string]string)

	for _, file := range getEnvVarsLocationsForZsh(env) {
		loadEnvVarsFromUser(file, envVars)
	}

	env.EnvVariables = envVars
}

func UnloadEnvVarsToUser(env *Env) {
	fileName := getMainEnvVarFile(env)

	for k, v := range env.EnvVariables {
		err := os.WriteFile(fileName, []byte("export "+k+"="+v), 0644)
		Check(err)
	}
}

func getMainEnvVarFile(env *Env) string {
	return env.Home() + "/.zshenv"
}

func getEnvVarsLocationsForZsh(env *Env) []string {
	return []string{
		getMainEnvVarFile(env),
		"/etc/zprofile",
		env.Home() + "/.zprofile",
		"/etc/zshrc",
		env.Home() + "/.zshrc",
		"/etc/zlogin",
		env.Home() + "/.zlogin",
		"/etc/zshenv",
	}
}

func loadEnvVarsFromUser(source string, envVars map[string]string) {
	_, err := os.Stat(source)

	if err != nil {
		return
	}

	lines := ParseFile(source)

	for _, line := range lines {
		split := strings.Split(line, " ")
		if len(split) != 2 {
			continue
		}

		if split[0] != "export" {
			continue
		}

		split2 := strings.Split(split[1], "=")

		if len(split) != 2 {
			continue
		}

		envVars[split2[0]] = split2[1]
	}
}