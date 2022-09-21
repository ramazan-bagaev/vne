package main

import (
	"os"
	"strings"
)

func GetUserEnvVars(env *Env) map[string]string {
	envVars := make(map[string]string)

	for _, file := range env.Shell.GetEnvVarsLocations(env) {
		loadEnvVarsFromUser(file, envVars)
	}

	return envVars
}

func UnloadEnvVarsToUser(env *Env, conf *Config) {
	fileName := env.Shell.GetMainEnvVarFile(env)

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	Check(err)
	defer f.Close()

	addedVars := ""

	for k, v := range conf.GetEnvVars() {
		addedVars += "\nexport " + k + "=" + v + "\n"
	}

	_, err = f.WriteString(addedVars)
	Check(err)
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
