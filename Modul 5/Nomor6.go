package main

import "fmt"

func power(x int, y int) int {
	if y == 0 {
		return 1
	}
	return x * power(x, y-1)
}

func main() {
	var x, y int
	fmt.Print("Masukkan bilangan x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan y: ")
	fmt.Scan(&y)
	
	result := power(x, y)
	fmt.Printf("Hasil %d pangkat %d adalah %d\n", x, y, result)
}
