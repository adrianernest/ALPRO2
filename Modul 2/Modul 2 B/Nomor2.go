package main

import "fmt"

func main() {
	var bunga, pita string
	var count int

	for {
		fmt.Printf("Bunga %d: ", count+1)
		fmt.Scan(&bunga)

		if bunga == "SELESAI" {
			break
		}

		if pita == "" {
			pita = bunga
		} else {
			pita += " - " + bunga
		}

		count++
	}

	fmt.Println("Pita:", pita+" -")
	fmt.Println("Bunga:",count)
}