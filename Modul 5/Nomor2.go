package main

import (
	"fmt"
)

func printStars(n int) {
	if n > 0 {
		fmt.Print("*")
		printStars(n - 1)
	}
}

func printPattern(n, i int) {
	if i <= n {
		printStars(i)
		fmt.Println()
		printPattern(n, i+1)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scan(&n)
	printPattern(n, 1)
}
