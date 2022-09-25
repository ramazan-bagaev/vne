package main

import (
	"log"
	"regexp"
	"strings"
)

type Config struct {
	configMap map[string][]string
}

func CreateEmptyConfig() *Config {
	conf := Config{}

	conf.configMap = make(map[string][]string)

	return &conf
}

func CreateConfig(path string) *Config {
	conf := Config{}

	configLines := ParseFile(path)
	configs := make(map[string][]string)
	currentConfName := "undefined"
	for _, line := range configLines {
		if match, _ := regexp.MatchString("\\[.", line); match {
			currentConfName = line[1 : len(line)-1]
		} else {
			configs[currentConfName] = append(configs[currentConfName], line)
		}
	}

	conf.configMap = configs

	return &conf
}

func (c *Config) GetEnvVars() map[string]string {
	envVars := make(map[string]string)
	for _, line := range c.configMap["envs"] {
		split := strings.Split(line, "=")
		if len(split) != 2 {
			log.Println("can't parse env variable")
			panic(1)
		}

		envVars[split[0]] = split[1]
	}

	return envVars
}

func (c *Config) GetTools() []string {
	return c.configMap["tools"]
}

func (c *Config) GetDirs() []string {
	return c.configMap["dirs"]
}

func (c *Config) GetShell() string {
	shell := c.configMap["shell"]

	if len(shell) != 1 {
		panic("shell should be specified and should be only one")
	}

	return c.configMap["shell"][0]
}

func (c *Config) SetEnvVars(varMap map[string]string) {
	vars := []string{}

	for k, v := range varMap {
		vars = append(vars, k+"="+v)
	}

	c.configMap["envVars"] = vars
}

func (c *Config) SetTools(tools []string) {
	c.configMap["tools"] = tools
}

func (c *Config) SetDirs(dirs []string) {
	c.configMap["dirs"] = dirs
}

func (c *Config) SetShell(shell string) {
	c.configMap["shell"] = []string{shell}
}

func (a *Config) Minus(b *Config) *Config {
	res := CreateEmptyConfig()

	aVars := a.GetEnvVars()
	bVars := b.GetEnvVars()

	resVars := make(map[string]string)
	for e1K, e1V := range aVars {
		if _, ok := bVars[e1K]; !ok {
			resVars[e1K] = e1V
		}
	}

	aTools := a.GetTools()
	bTools := b.GetTools()

	resTools := []string{}
	for _, e1T := range aTools {
		uniq := true
		for _, e2T := range bTools {
			if e1T == e2T {
				uniq = false
			}
		}

		if uniq {
			resTools = append(resTools, e1T)
		}
	}

	res.SetEnvVars(resVars)
	res.SetTools(resTools)
	res.SetDirs(a.GetDirs())
	res.SetShell(a.GetShell())

	return res
}
