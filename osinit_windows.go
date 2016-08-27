package main

import (
  "os/exec"
  "log"
)

func init_os(){
  setTermSize := exec.Command("mode", "160,60")
  err := setTermSize.Run()
  if err != nil {
    log.Println(err)
  }
}
