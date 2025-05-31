package main

import (
	"log"
	"os"
	"os/exec"
	"strconv"
)

var appSet_1 = []string{
	"/Applications/Spotify.app",
	"/Applications/Safari.app",
}

var appSet_2 = []string{
	"/Applications/Discord.app",
	"/Applications/Google.app",
}

func main() {

	if len(os.Args) != 2 {
		log.Fatal("wrong Args num ")
	}

	ar, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("ATOI ERR")
	}

	if ar == 2 {
		for _, path := range appSet_1 {
			cmd := exec.Command("open", path)
			cmd.Run()
		}
	} else if ar == 3 {
		for _, path := range appSet_2 {
			cmd := exec.Command("open", path)
			cmd.Run()
		}
	} else {
		log.Fatal("NaN")
	}
}