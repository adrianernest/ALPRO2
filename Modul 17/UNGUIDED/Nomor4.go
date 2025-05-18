package main

import (
	"fmt"
	"math"
)

func pembulatan(x float64) float64 {
	return math.Trunc(x*1e10) / 1e10
}

func hitungPI(n int) (float64, float64, int) {
	total, sebelumnya := 0.0, 0.0

	for i, tanda := 0, 1.0; i < n; i, tanda = i+1, -tanda {
		total += tanda / float64(2*i+1)
		sekarang := total * 4

		if i > 0 && math.Abs(sekarang-sebelumnya) < 0.00001 {
			return pembulatan(sebelumnya), pembulatan(sekarang), i + 1
		}
		sebelumnya = sekarang
	}
	return pembulatan(sebelumnya), pembulatan(sebelumnya), n
}

func main() {
	var suku int
	fmt.Print("N suku pertama: ")
	fmt.Scan(&suku)

	pi1, pi2, iter := hitungPI(suku)

	fmt.Printf("Hasil PI: %.10f\n", pi1)
	fmt.Printf("Hasil PI: %.10f\n", pi2)
	fmt.Printf("Pada i ke: %d\n", iter)
}
