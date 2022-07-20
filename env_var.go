package main

import(
	"os"
	"strings"
)

func LoadEnvVarsFromUser(env *Env) {
	conf := env.ConfigPath()

	envVars := make(map[string]string)

	for _, file := range getEnvVarsLocationsForZsh(env) {
		loadEnvVars(file, envVars)
	}

	env.EnvVariables = envVars
	updateConfig(conf, &envVars)
}

func getEnvVarsLocationsForZsh(env *Env) []string {
	return []string {
		env.Home() + "/.zshenv",
		"/etc/zprofile",
		env.Home() + "/.zprofile",
		"/etc/zshrc",
		env.Home() + "/.zshrc",
		"/etc/zlogin",
		env.Home() + "/.zlogin",
		"/etc/zshenv",
	}
}

func loadEnvVars(source string, envVars map[string]string) {
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

func updateConfig(config string, envVars *map[string]string) {

}