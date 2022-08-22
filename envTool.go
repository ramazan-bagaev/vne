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

func LoadToolsToVNE(env *Env) {
	tools := []string{}

	ct := GetCreationDate()

	for _, dir := range getEnvToolsLocationsForZsh(env) {
		tools = append(tools, readFromBinDir(dir, ct)...)
	}

	env.Tools = tools
}

func getEnvToolsLocationsForZsh(env *Env) []string {
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

func readFromBinDir(dir string, ct time.Time) []string {
	bins, err := ioutil.ReadDir(dir)

	if err != nil {
		return []string{}
	}

	binFiles := []string{}

	for _, bin := range bins {
		if bin.ModTime().Equal(ct) {
			continue
		}

		binFiles = append(binFiles, bin.Name())
	}

	return binFiles
}
