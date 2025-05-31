package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const maxNum = 1

var appSetNew []string

func main() {
	var t string

	if len(os.Args) != 1 {
		log.Fatal("wrong Args num ")
	} else {
		fmt.Println("Add Your set max", maxNum)
		for i := 0; i < maxNum; i++ {
			_, err := fmt.Scan(&t)
			if err != nil {
				log.Fatal(err)
			}
			appSetNew = append(appSetNew, t)
		}

		data := []byte(strings.Join(appSetNew, "\n"))
		err := os.WriteFile("appSet.txt", data, 0777)
		if err != nil {
			fmt.Println(err)
		}

		cmd := exec.Command("go", "build", "executer.go")
		cmd.Run()
		

	}

}
