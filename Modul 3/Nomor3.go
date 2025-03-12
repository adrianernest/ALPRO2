package main

import "fmt"

func jarak(a, b, c, d float64) float64 {
	return (a-c)*(a-c) + (b-d)*(b-d)
}

func didalam(cx, cy, r, x, y float64) int {
	if jarak(cx, cy, x, y) <= r*r {
		return 1
	}
	return 0
}

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y float64
	fmt.Scan(&cx1, &cy1, &r1, &cx2, &cy2, &r2, &x, &y)

	pos := didalam(cx1, cy1, r1, x, y)<<1 | didalam(cx2, cy2, r2, x, y)

	switch pos {
	case 3:
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	case 2:
		fmt.Println("Titik di dalam lingkaran 1")
	case 1:
		fmt.Println("Titik di dalam lingkaran 2")
	default:
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
