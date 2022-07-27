package main

import (
	"io/ioutil"
	"log"
)

func GetTools(configs map[string][]string) []string {
	return configs["tools"]
}

func LoadToolsToVNE(env *Env) {
	tools := []string{}

	for _, dir := range getEnvToolsLocationsForZsh(env) {
		tools = append(tools, readFromBinDir(dir)...)
	}

	env.Tools = tools
}

func getEnvToolsLocationsForZsh(env *Env) []string {
	return []string{
		"/usr/local/bin",
		"/usr/bin",
		"/bin",
		"/usr/sbin",
		"/sbin"}
}

func readFromBinDir(dir string) []string {
	bins, err := ioutil.ReadDir(dir)

	if err != nil {
		return []string{}
	}

	binFiles := []string{}

	for _, bin := range bins {
		log.Printf("bin name %s", bin.Name())
		binFiles = append(binFiles, bin.Name())
	}

	return binFiles
}
