package main

import (
	"os/exec"
)

var paths = []string{
	"/Applications/Spotify.app",
	"/Applications/Safari.app",
}

func main() {

	for _, path := range paths {
		cmd := exec.Command("open", path)
		cmd.Run()
	}

}