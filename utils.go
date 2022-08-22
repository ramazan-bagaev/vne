package main

import (
	"bufio"
	"os"
)

func ParseFile(path string) []string {
	file, err := os.Open(path)
	Check(err)
	defer file.Close()

	scn := bufio.NewScanner(file)

	res := make([]string, 0)
	for scn.Scan() {
		str := scn.Text()

		res = append(res, str)
	}

	return res
}

func Check(e error) {
	if e != nil {
		panic(e.Error())
	}
}
