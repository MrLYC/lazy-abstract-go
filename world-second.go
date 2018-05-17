// +build second

package main

import (
	"fmt"
)

// TheWorld :
type TheWorld struct{}

type World = *TheWorld

// Hello :
func (w *TheWorld) Hello() {
	fmt.Println("hello World")
}

// NewWorld :
func NewWorld() World {
	return &TheWorld{}
}
