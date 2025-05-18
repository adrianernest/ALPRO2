package main

import (
	"fmt"
)

func main() {
	var angka float64
	var total float64
	var jumlah int

	fmt.Scanln(&angka)

	if angka == 9999 {
		fmt.Println("Tidak ada angka yang dimasukan")
	} else {
		total = angka
		jumlah = 1

		for {
			fmt.Scanln(&angka)
			if angka == 9999 {
				break
			}
			total += angka
			jumlah++
		}

		rataRata := total / float64(jumlah)
		fmt.Printf("Rata-rata bilangan adalah: %.2f\n", rataRata)
	}
}
