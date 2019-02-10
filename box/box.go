package box

import (
	. "fmt"
)

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width float64
	height float64
	depth float64
	color Color
}

type BoxList []Box

func (box Box) Volume() (v float64) {
	v = box.width * box.height * box.depth

	return v
}

func (c Color) String() (colorName string) {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}

	return strings[c]
}

func (boxList BoxList) BiggestColor() (color Color) {
	maxVol := 0.0
	color = WHITE

	for _, box := range boxList {
		vol := box.Volume()

		if vol > maxVol {
			color = box.color
		}
	}

	return color
}

func (bl *BoxList) PaintItBlack() {
	for i, _ := range *bl {
		listEnt := *bl

		listEnt[i].color = BLACK
	}
}

func BoxImpl() {
	boxes := BoxList {
		{4, 4, 4, RED},
		{10, 10, 1, YELLOW},
		{1, 1, 20, BLACK},
		{10, 10, 1, BLUE},
		{10, 30, 1, WHITE},
		{20, 20, 20, YELLOW},
	}

	Printf("We have %d boxes in our set.\n", len(boxes))
	Println("The volume of the first one is", boxes[0].Volume(), "cm³.")
	Println("The color of the last one is", boxes[len(boxes) - 1].color.String())
	Println("The biggest one is", boxes.BiggestColor().String(), ".")

	Println("Let's paint them all black.")

	boxes.PaintItBlack()

	Println("The color of the second one is", boxes[1].color.String())
	Println("Obviously, the biggest one is", boxes.BiggestColor().String(), ".")
}
