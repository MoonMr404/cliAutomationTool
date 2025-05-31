package main

import (
	"os"
	"os/exec"
)

func main() {
	
	data, err := os.ReadFile("appSet.txt")
	if err != nil {
		panic(err)
	}
	//made this way in order to be able to read the path as string
	lines := string(data)
	for _, path := range []string{lines} {
		cmd := exec.Command("open", path)
		cmd.Run()
	}
}
