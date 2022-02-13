package main

import "fmt"

type Direction int

const (
	None Direction = iota
	North
	East
	South
	West
)

func DirectionToString(d Direction) string {
	switch d {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	default:
		return "None"
	}
}

func GetDirection(d float64) Direction {
	switch {
	case d >= 315 && d <= 360, d >= 0 && d < 45:
		return North
	case d >= 45 && d < 135:
		return East
	case d >= 135 && d < 225:
		return South
	case d >= 225 && d < 315:
		return West
	default:
		return None
	}
}

func main() {
	fmt.Println(DirectionToString(GetDirection(38.3)))
}
