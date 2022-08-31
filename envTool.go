package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func GetTools(configs map[string][]string) []string {
	return configs["tools"]
}

func GetUserTools(env *Env) []string {
	tools := []string{}

	ct := GetCreationDate()

	for _, dir := range getEnvToolsLocations(env) {
		tools = append(tools, readFromBinDirNotWithSameCT(dir, ct)...)
	}

	return tools
}

func getEnvToolsLocations(env *Env) []string {
	path := os.Getenv("PATH")
	return strings.Split(path, ":")
}

func GetCreationDate() time.Time {
	etcFiles, err := ioutil.ReadDir("/etc")

	if err != nil {
		log.Fatal("failed to find how old unix system is", err.Error())
	}

	ct := time.Now()

	for _, file := range etcFiles {
		t := file.ModTime()

		if ct.After(t) {
			ct = t
		}

	}

	return ct
}

func readFromBinDirNotWithSameCT(dir string, ct time.Time) []string {
	bins, err := ioutil.ReadDir(dir)

	if err != nil {
		return []string{}
	}

	binFiles := []string{}

	for _, bin := range bins {
		if strings.HasPrefix(bin.Name(), ".") {
			continue
		}

		if bin.ModTime().Equal(ct) {
			continue
		}

		binFiles = append(binFiles, bin.Name())
	}

	return binFiles
}

func GetAvailableToolsFromList(env *Env, tools []string) []string {
	availableTools := []string{}
	for _, dir := range getEnvToolsLocations(env) {
		bins := readFromBinDir(dir)

		for _, tool := range tools {
			for _, bin := range bins {
				if bin == tool {
					availableTools = append(availableTools, tool)
				}
			}
		}
	}

	return availableTools
}

func readFromBinDir(dir string) []string {
	bins, err := ioutil.ReadDir(dir)

	if err != nil {
		return []string{}
	}

	binFiles := []string{}

	for _, bin := range bins {
		binFiles = append(binFiles, bin.Name())
	}

	return binFiles
}
