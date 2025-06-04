package main

import (
	"os"
	"os/exec"
	"strings"
)

func main() {

	data, err := os.ReadFile("C:\\Users\\Francesco\\Desktop\\appSet.txt")
	if err != nil {
		panic(err)
	}
	paths := strings.Split(string(data), "\n")
	for _, path := range paths {
		if path == "" {
			continue
		}
		cmd := exec.Command("cmd", "/C", "start", "", path)
		cmd.Run()
	}
}
