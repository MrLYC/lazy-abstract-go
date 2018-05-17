// +build first

package main

import (
	"fmt"
)

// World :
type World struct{}

// Hello :
func (w *World) Hello() {
	fmt.Println("hello world")
}

// NewWorld :
func NewWorld() World {
	return World{}
}
