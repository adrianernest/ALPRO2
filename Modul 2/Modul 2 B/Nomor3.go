package main

import "fmt"

func main() {
	var berat1, berat2 float64
	totalBerat := 0.0

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&berat1, &berat2)

		if berat1 < 0 || berat2 < 0 || (totalBerat+berat1+berat2) > 150 {
			break
		}

		selisih := berat1 - berat2
		if selisih < 0 {
			selisih = -selisih
		}

		fmt.Println("Sepeda motor pak Andi akan oleng:", selisih >= 9)

		totalBerat += berat1 + berat2
	}

	fmt.Println("Proses selesai.")
}
