package main

import (
	"os"
	"strings"
)

func LoadEnvVarsToVNE(env *Env) {
	conf := env.ConfigPath()

	envVars := make(map[string]string)

	for _, file := range getEnvVarsLocationsForZsh(env) {
		loadEnvVarsFromUser(file, envVars)
	}

	env.EnvVariables = envVars
	updateVNEConfig(conf, &envVars)
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

func updateVNEConfig(config string, envVars *map[string]string) {
	content := "[envs]\n"

	for key, value := range *envVars {
		content += key + "=" + value + "\n"
	}

	err := os.WriteFile(config, []byte(content), 0644)
	Check(err)
}
