package main

import (
	"fmt"
)

var startDate = Date{102, 4, 19}

var months = []Month{
	Month{30, "Wormkiln"},  //Roughly January
	Month{30, "Wormlax"},   //Roughly Febuary
	Month{30, "Freshstar"}, //Roughly March
	Month{29, "Newlife"},   //Roughly April
	Month{30, "Lifeborn"},  //Roughly May
	Month{29, "Burnstar"},  //Roughly June
	Month{30, "Brightsee"}, //Roughly July
	Month{30, "Lightlax"},  //Roughly August
	Month{29, "Darkstar"},  //Roughly September
	Month{29, "Deathsky"},  //Roughly October
	Month{30, "Laxday"},    //Roughly November
	Month{29, "Wormstar"},  //Roughly December
}

type Date struct {
	year  uint16
	month uint8
	day   uint8
}

type Month struct {
	days uint8
	name string
}

func (d Date) String() string {
	return fmt.Sprintf("%d %s %d PB", d.day, months[d.month].String(), d.year)
}

func (m Month) String() string {
	return m.name
}

func (d Date) Next() Date {
	nextDate := d
	m := months[d.month]
	if d.day == m.days {
		nextDate.day = 1
		if d.month == 11 {
			nextDate.month = 0
			nextDate.year++
		} else {
			nextDate.month++
		}
	} else {
		nextDate.day++
	}
	return nextDate
}
