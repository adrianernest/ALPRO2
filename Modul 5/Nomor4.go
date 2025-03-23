package main

import "fmt"

func printDescending(n int) {
	if n < 1 {
		return
	}
	fmt.Printf("%d ", n)
	printDescending(n - 1)
}

func printAscending(n, current int) {
	if current > n {
		return
	}
	fmt.Printf("%d ", current)
	printAscending(n, current+1)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan positif: ")
	fmt.Scan(&n)
	printDescending(n)
	printAscending(n, 2)
	fmt.Println()
}
