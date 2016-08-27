package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
  "strings"
)

const (
  creationStatePlayerName = iota
  creationStatePlayerProfession
  creationStatePartyMembers
  creationStateConfirm
)

const (
  initCreationInterfaceXOffset = 12
  initCreationInterfaceYUpOffset = 1
  initCreationInterfaceYDownOffset = 1
)

const partyCount = 5

var creationState int

var (
  creationInterfaceXOffset int
  creationInterfaceYUpOffset int
  creationInterfaceYDownOffset int
)

func initCreationScreenSettings() {
  creationInterfaceXOffset = initCreationInterfaceXOffset
  creationInterfaceYUpOffset = initCreationInterfaceYUpOffset
  creationInterfaceYDownOffset = initCreationInterfaceYDownOffset
}

func charCreationLayout(g *gocui.Gui) error {
  maxX, maxY := g.Size()
  if v, err := g.SetView("charBackground", 0, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
    v.Title = "Character Creation"
	}
  if v, err := g.SetView("charInfo", 0, maxY*3/4, maxX-1, maxY-1); err != nil {
    if err != gocui.ErrUnknownView {
      return err
    }
    v.Title = "Information"
    fmt.Fprintln(v, "Please enter the name for your character")
  }
  if v, err := g.SetView("creationInterface", maxX/2-creationInterfaceXOffset,
      maxY/2-creationInterfaceYUpOffset,
      maxX/2+creationInterfaceXOffset,
      maxY/2+creationInterfaceYDownOffset); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
    creationState = creationStatePlayerName
    g.Cursor = true
    g.Editor = gocui.EditorFunc(fieldEditor)
    v.Title = "What is Your Name?"
    v.Editable = true
    if err := g.SetCurrentView("creationInterface"); err != nil {
			return err
    }
	}
  return nil
}

func creationSelect(g *gocui.Gui, v *gocui.View) error {
  background, err := g.View("charBackground")
  if err != nil {
    return err
  }
  info, err := g.View("charInfo")
  if err != nil {
    return err
  }
  switch creationState {
  case creationStatePlayerName:
    str, err := v.Line(0)
    if err != nil {
      str = ""
    }
    str = strings.TrimSpace(str)
    if len(str) == 0 {
      info.Clear()
      info.SetOrigin(0, 0)
      fmt.Fprintln(info, "You MUST provide a name for your character")
      return nil
    }
    WorldState.player.character.name = str
    fmt.Fprintf(background, "Player's Name: %s\n", str)
    v.Clear()
    v.SetCursor(0, 0)
    v.SetOrigin(0, 0)
    creationInterfaceXOffset = 15
    creationInterfaceYUpOffset = professionListOffset + 1
    creationInterfaceYDownOffset = professionListOffset + professionListOddness
    g.Cursor = false
    v.Editable = false
    v.Highlight = true
    v.Title = "Select your Profession"
    for _, p := range(professionList) {
      fmt.Fprintln(v, p.name)
    }
    creationState = creationStatePlayerProfession
    info.Clear()
    info.SetOrigin(0, 0)
    fmt.Fprintln(info, professionList[0].description)
  case creationStatePlayerProfession:
    _, cy := v.Cursor()
    WorldState.player.profession = professionList[cy]
    fmt.Fprintf(background, "Player's Profession: %s\n", WorldState.player.profession.name)
    v.Clear()
    v.SetCursor(0, 0)
    v.SetOrigin(0, 0)
    creationInterfaceXOffset = initCreationInterfaceXOffset
    creationInterfaceYUpOffset = initCreationInterfaceYUpOffset
    creationInterfaceYDownOffset = initCreationInterfaceYDownOffset
    g.Cursor = true
    v.Editable = true
    v.Highlight = false
    creationState = creationStatePartyMembers
    v.Title = "Party Member Name"
    info.Clear()
    info.SetOrigin(0, 0)
    fmt.Fprintf(info,"Name your Party Member (1/%d)", partyCount)
  case creationStatePartyMembers:
    str, err := v.Line(0)
    if err != nil {
      str = ""
    }
    str = strings.TrimSpace(str)
    if len(str) == 0 {
      info.Clear()
      info.SetOrigin(0, 0)
      fmt.Fprintf(info,"Name your Party Member (%d/%d)\n", len(WorldState.player.party) + 1, partyCount)
      fmt.Fprintln(info, "You MUST provide a name for your party members")
      return nil
    }
    c := new(Character)
    c.name = str
    WorldState.player.party = append(WorldState.player.party, c)
    fmt.Fprintf(background, "Party Member %d/%d: %s\n", len(WorldState.player.party), partyCount, str)
    v.Clear()
    v.SetOrigin(0, 0)
    v.SetCursor(0, 0)
    if len(WorldState.player.party) == partyCount {
      g.Cursor = false
      v.Editable = false
      v.Highlight = true
      v.Title = "Ready?"
      creationInterfaceXOffset = 5
      info.Clear()
      info.SetOrigin(0, 0)
      fmt.Fprintln(info,"Press Enter to start your Adventure!")
      fmt.Fprint(v, "Yes!")
      creationState = creationStateConfirm
    } else {
      info.Clear()
      info.SetOrigin(0, 0)
      fmt.Fprintf(info,"Name your Party Member (%d/%d)", len(WorldState.player.party) + 1, partyCount)
    }
  case creationStateConfirm:
    WorldState.player.Init()
    g.DeleteView("charBackground")
    g.DeleteView("charInfo")
    g.DeleteView("charInterface")
    g.SetLayout(layout)
  }
  return nil
}

func creationUp(g *gocui.Gui, v *gocui.View) error {
  info, err := g.View("charInfo")
  if err != nil {
    return err
  }
  switch creationState {
  case creationStatePlayerProfession:
    if v != nil {
  		cx, cy := v.Cursor()
  		if cy > 0 {
        cy--
  			if err := v.SetCursor(cx, cy); err != nil {
          return err
        }
        info.Clear()
        info.SetOrigin(0, 0)
        fmt.Fprintln(info, professionList[cy].description)
  		}
  	}
  }
  return nil
}

func creationDown(g *gocui.Gui, v *gocui.View) error {
  info, err := g.View("charInfo")
  if err != nil {
    return err
  }
  switch creationState {
  case creationStatePlayerProfession:
    if v != nil {
  		cx, cy := v.Cursor()
  		if cy < len(professionList) - 1 {
        cy++
  			if err := v.SetCursor(cx, cy); err != nil {
          return err
        }
        info.Clear()
        info.SetOrigin(0, 0)
        fmt.Fprintln(info, professionList[cy].description)
  		}
  	}
  }
  return nil
}
