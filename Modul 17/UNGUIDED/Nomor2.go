package main

import (
	"fmt"
)

func cariString(target string, banyak int) {
	var posisi []int

	for i := 1; i <= banyak; i++ {
		var input string
		fmt.Print("Masukan kata ke-", i, ": ")
		fmt.Scanln(&input)
		if input == target {
			posisi = append(posisi, i)
		}
	}

	if len(posisi) > 0 {
		fmt.Print("Kata ditemukan pada posisi: ")
		for i, p := range posisi {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(p)
		}
		fmt.Println("\nJumlah kemunculan:", len(posisi))
		if len(posisi) >= 2 {
			fmt.Println("Ada sedikitnya dua kata", target)
		} else {
			fmt.Println("Hanya ada satu string", target)
		}
	} else {
		fmt.Println("Kata tidak ditemukan.")
	}
}

func main() {
	var target string
	var banyak int

	fmt.Print("Masukan kata yang akan dicari: ")
	fmt.Scanln(&target)
	fmt.Print("Masukan banyak data: ")
	fmt.Scanln(&banyak)

	cariString(target, banyak)
}
