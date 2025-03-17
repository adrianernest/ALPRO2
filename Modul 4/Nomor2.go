package main

import (
	"fmt"
)

func hitungSkor(waktu [8]int) (int, int) {
	const batasWaktu = 301
	soal, total := 0, 0

	for _, t := range waktu {
		if t < batasWaktu {
			soal++
			total += t
		}
	}
	return soal, total
}

func main() {
	var nama, pemenang string
	var waktu [8]int
	maxSoal, minWaktu := 0, 0

	for {
		fmt.Scan(&nama)
		if nama == "selesai" || nama == "Selesai" {
			break
		}

		for i := range waktu {
			fmt.Scan(&waktu[i])
		}

		soal, total := hitungSkor(waktu)

		if soal > maxSoal || (soal == maxSoal && total < minWaktu) {
			pemenang, maxSoal, minWaktu = nama, soal, total
		}
	}

	fmt.Println(pemenang, maxSoal, minWaktu)
}
