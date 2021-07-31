package main

import (
	"fmt"
)

type color int

const (
	R color = iota
	G
	B
	W
	Y
	O
)

type Cube = [6][3][3]color

var c = Cube{
	{{R, R, R},
		{R, R, R},
		{R, R, R}},
	{{G, G, G},
		{G, G, G},
		{G, G, G}},
	{{B, B, B},
		{B, B, B},
		{B, B, B}},
	{{W, W, W},
		{W, W, W},
		{W, W, W}},
	{{Y, Y, Y},
		{Y, Y, Y},
		{Y, Y, Y}},
	{{O, O, O},
		{O, O, O},
		{O, O, O}},
}

func main() {
	fmt.Println(c)
}
