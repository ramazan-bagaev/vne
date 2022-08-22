package main

import (
	"strings"
)

func GetConsole(shellPath string) Console {
	spl := strings.Split(shellPath, "/")

	cnsl := spl[len(spl)-1]

	if cnsl == "zsh" {
		return Zsh{}
	}

	panic("can't recognize console: " + cnsl)
}

type Zsh struct {
}

type Console interface {
	GetEnvVarsLocations(env *Env) []string
	GetMainEnvVarFile(env *Env) string
	Name() string
}

func (c Zsh) GetEnvVarsLocations(env *Env) []string {
	return []string{
		c.GetMainEnvVarFile(env),
		"/etc/zprofile",
		env.Home() + "/.zprofile",
		"/etc/zshrc",
		env.Home() + "/.zshrc",
		"/etc/zlogin",
		env.Home() + "/.zlogin",
		"/etc/zshenv",
	}
}

func (c Zsh) GetMainEnvVarFile(env *Env) string {
	return env.Home() + "/.zshenv"
}

func (c Zsh) Name() string {
	return "zsh"
}
