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

func (c color) String() string {
	var lookup = [6]string{"R", "G", "B", "W", "Y", "O"}
	return lookup[c]
}

/*
A matrix of faces, representing a cube. Each face is a 3x3 matrix of colour
values

We number the faces of the cube like so:

    0
  1 2 3
    4
    5

*/
type Cube [6][3][3]color

// Formats the cube as a net with the faces numbered as above
func (c Cube) String() string {
	var res string

	res += "\n"
	formatFaceInCenter(&res, c[0])

	row2 := c[1:4]
	for i := 0; i < 3; i++ {
		for _, face := range row2 {
			for _, color := range face[i] {
				res += " "
				res += color.String()
				res += " "
			}
			res += " "
		}
		res += "\n"
	}
	res += "\n"

	formatFaceInCenter(&res, c[4])
	formatFaceInCenter(&res, c[5])

	return res
}

func formatFaceInCenter(res *string, face [3][3]color) {
	for _, row := range face {
		*res += "          "
		for _, color := range row {
			*res += " "
			*res += color.String()
			*res += " "
		}
		*res += "\n"
	}
	*res += "\n"
}

var c = Cube{
	{{R, G, B},
		{R, G, B},
		{R, G, B}},
	{{R, G, B},
		{R, G, B},
		{R, G, B}},
	{{R, G, B},
		{R, G, B},
		{R, G, B}},
	{{R, G, B},
		{R, G, B},
		{R, G, B}},
	{{Y, Y, Y},
		{Y, Y, Y},
		{Y, Y, Y}},
	{{O, O, O},
		{O, O, O},
		{O, O, O}},
}

// var c = Cube{
// 	{{R, R, R},
// 		{R, R, R},
// 		{R, R, R}},
// 	{{G, G, G},
// 		{G, G, G},
// 		{G, G, G}},
// 	{{B, B, B},
// 		{B, B, B},
// 		{B, B, B}},
// 	{{W, W, W},
// 		{W, W, W},
// 		{W, W, W}},
// 	{{Y, Y, Y},
// 		{Y, Y, Y},
// 		{Y, Y, Y}},
// 	{{O, O, O},
// 		{O, O, O},
// 		{O, O, O}},
// }

func main() {
	fmt.Println(c)
}
