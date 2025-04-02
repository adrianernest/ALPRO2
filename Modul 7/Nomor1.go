package main

import (
	"fmt"
)

type Titik struct {
	x, y int
}

type Lingkaran struct {
	pusat  Titik
	radius int
}

func dalamLingkaran(c Lingkaran, p Titik) bool {
	dx, dy := p.x-c.pusat.x, p.y-c.pusat.y
	return dx*dx+dy*dy <= c.radius*c.radius
}

func main() {
	var c1, c2 Lingkaran
	var t Titik

	fmt.Scan(&c1.pusat.x, &c1.pusat.y, &c1.radius)
	fmt.Scan(&c2.pusat.x, &c2.pusat.y, &c2.radius)
	fmt.Scan(&t.x, &t.y)

	in1, in2 := dalamLingkaran(c1, t), dalamLingkaran(c2, t)

	switch {
	case in1 && in2:
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	case in1:
		fmt.Println("Titik di dalam lingkaran 1")
	case in2:
		fmt.Println("Titik di dalam lingkaran 2")
	default:
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
