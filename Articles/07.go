package main

import (
	"fmt"
	"strings"
	"time"
)

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type rover struct{}

func (r rover) talk() string {
	return "whir whir"
}

// // stardate возвращает выдуманное измерение времени для указанной даты.
// func stardate(t time.Time) float64 {
// 	doy := float64(t.YearDay())
// 	h := float64(t.Hour()) / 24.0
// 	return 1000 + doy + h
// }

type stardater interface {
	YearDay() int
	Hour() int
}

// stardate возвращает выдуманное измерение времени.
func stardate(t stardater) float64 {
	doy := float64(t.YearDay())
	h := float64(t.Hour()) / 24.0
	return 1000 + doy + h
}

type sol int

func (s sol) YearDay() int {
	return int(s % 668) // Марсианской год состоит из 668 дней
}

func (s sol) Hour() int {
	return 0 // Неизвестный час
}

// location с широтой и долготой в десятичных градусах
type location struct {
	lat, long coordinate
}

// String форматирует location с широтой и долготой
func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.lat, l.long)
}

type coordinate struct {
	d, m, s float64
	h       rune
}

func (c coordinate) String() string {
	return fmt.Sprintf("%v°%v'%.1f\" %c", c.d, c.m, c.s, c.h)
}

// Elysium Planitia is at 4°30'0.0" N, 135°54'0.0" E

func main() {
	var t interface {
		talk() string
	}

	t = martian{}
	fmt.Println(t.talk()) // Выводит: nack nack

	t = laser(3)
	fmt.Println(t.talk()) // Выводит: pew pew pewvv

	shout(martian{}) // Выводит: NACK NACK
	shout(laser(2))  // Выводит: PEW PEW

	t = rover{}
	shout(t) // Выводит: WHIR WHIR

	day := time.Date(2012, 8, 6, 5, 17, 0, 0, time.UTC)
	fmt.Printf("%.1f Curiosity has landed\n", stardate(day)) // Выводит: 1219.2 Curiosity has landed

	s := sol(1422)
	fmt.Printf("%.1f Happy birthday\n", stardate(s)) // Выводит: 1086.0 Happy birthday

	// curiosity := location{-4.5895, 137.4417}
	// fmt.Println(curiosity) // Выводит: -4.5895, 137.4417

	elysium := location{
		lat:  coordinate{4, 30, 0.0, 'N'},
		long: coordinate{135, 54, 0.0, 'E'},
	}
	fmt.Println("Elysium Planitia is at", elysium) // Выводит: Elysium Planitia is at 4°30’0.0” N, 135°54’0.0” E

}
