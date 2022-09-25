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

	if shl == "bash" {
		return Bash{}
	}

	panic("can't recognize shell: " + shl)
}

type Shell interface {
	GetEnvVarsLocations(env *Env) []string
	GetMainEnvVarFile(env *Env) string
	Name() string
	Path() string
}

// zsh

type Zsh struct {
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

func (c Zsh) Path() string {
	return "/bin/zsh"
}

// bash

type Bash struct {
}

func (b Bash) GetEnvVarsLocations(env *Env) []string {
	return []string{
		"/etc/profile",
		b.GetMainEnvVarFile(env),
		env.Home() + "/.bash_login",
		env.Home() + "/.profile",
		env.Home() + "/.bashrc",
	}
}

func (b Bash) GetMainEnvVarFile(env *Env) string {
	return env.Home() + "/.bash_profile"
}

func (b Bash) Name() string {
	return "bash"
}

func (b Bash) Path() string {
	return "/bin/bash"
}
