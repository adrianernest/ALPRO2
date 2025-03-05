package main

import "fmt"

func main() {
	var berat int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&berat)

	kg := berat / 1000
	gram := berat % 1000

	biayaKg := kg * 10000
	biayaGram := 0

	if gram > 0 {
		if gram >= 500 {
			biayaGram = gram * 5
		} else {
			biayaGram = gram * 15
		}
	}

	if kg >= 10 {
		biayaGram = 0
	}

	total := biayaKg + biayaGram

	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, gram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaGram)
	fmt.Printf("Total biaya: Rp. %d\n", total)
}
