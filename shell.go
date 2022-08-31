package main

import (
	"strings"
)

func GetShell(shellPath string) Shell {
	spl := strings.Split(shellPath, "/")

	shl := spl[len(spl)-1]

	if shl == "zsh" {
		return Zsh{}
	}

	panic("can't recognize shell: " + shl)
}

type Zsh struct {
}

type Shell interface {
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
