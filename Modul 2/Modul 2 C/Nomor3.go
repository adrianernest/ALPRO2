package main

import "fmt"

func main() {
	var b int
	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	fmt.Print("Faktor: ")
	count := 0
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i, " ")
			count++
		}
	}
	fmt.Println()

	if count == 2 {
		fmt.Println("Prima: true")
	} else {
		fmt.Println("Prima: false")
	}
}
