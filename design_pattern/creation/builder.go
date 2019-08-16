package main

import (
	"fmt"
)

/*
	Builder is a creational design pattern that lets you construct complex objects step by step.
	The pattern allows you to produce different types and representations of an object using the same construction code.
	https://refactoring.guru/design-patterns/builder
*/

type house struct {
	window int
	door   int
	garden bool
}

// HouseBuilder new instance object
func HouseBuilder() *house {
	return &house{}
}

func (h *house) setWindow(w int) *house {
	h.window = w
	return h
}

func (h *house) setDoor(d int) *house {
	h.door = d
	return h
}

func (h *house) setGarden(g bool) *house {
	h.garden = g
	return h
}

func main() {
	golfHouse := HouseBuilder().setDoor(4).setWindow(4).setGarden(true)

	fmt.Println(golfHouse)
}
