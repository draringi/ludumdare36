package main

import (
  "github.com/nsf/termbox-go"
  "log"
)

func main() {
  init_os()
  err :=  termbox.Init()
  if err != nil {
    log.Println(err)
    return
  }
  defer termbox.Close()
  mainloop:
  for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break mainloop
			}
		case termbox.EventError:
			log.Println(ev.Err)
      return
		}
}
}
