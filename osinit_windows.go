package main

import (
	"log"
	"os/exec"
)

func init_os() {
	//TODO: Change this to something less fragile
	setTermSize := exec.Command("mode", "160,60")
	err := setTermSize.Run()
	if err != nil {
		log.Panicln(err)
	}
}
