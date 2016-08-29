package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

type LogEntry struct {
	Date Date
	Msg  string
}

func (l LogEntry) String() string {
	return fmt.Sprintf("[%v] %s", l.Date, l.Msg)
}

func printLogsToView(logs []LogEntry, v *gocui.View) {
	v.Clear()
	v.SetOrigin(0, 0)
	_, height := v.Size()
	var sliceStart int
	if len(logs) > height {
		sliceStart = len(logs) - height
	}
	for _, l := range logs[sliceStart:] {
		fmt.Fprintln(v, l)
	}
}

func NewLogEntry(date Date, formatString string, v ...interface{}) LogEntry {
	return LogEntry{date, fmt.Sprintf(formatString, v...)}
}
