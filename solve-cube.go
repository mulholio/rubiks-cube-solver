package main

import (
	"fmt"
	"strconv"
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

/*
A matrix of faces, representing a cube. Each face is a 3x3 matrix of colour
values

We number the faces of the cube like so:

    1
  2 3 4
    5
    6

*/
type Cube [6][3][3]color

func (c Cube) String() string {
	var res string
	for _, face := range c {
		for _, row := range face {
			for _, color := range row {
				res += strconv.Itoa(int(color))
				res += " "
			}
			res += "\n"
		}
		res += "\n"
	}
	return res
}

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
