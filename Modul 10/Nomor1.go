package main

import (
	"fmt"
)

func main() {
	var N int

	// Input jumlah anak kelinci
	fmt.Print("Masukkan jumlah anak kelinci: ")
	fmt.Scan(&N)

	// Validasi jumlah anak kelinci
	if N <= 0 || N > 1000 {
		fmt.Println("Jumlah anak kelinci harus antara 1 hingga 1000.")
		return
	}

	// Input berat anak kelinci
	weights := make([]float64, N)
	fmt.Println("Masukkan berat anak kelinci:")
	for i := 0; i < N; i++ {
		fmt.Scan(&weights[i])
	}

	// Inisialisasi nilai minimum dan maksimum
	minWeight := weights[0]
	maxWeight := weights[0]

	// Mencari berat terkecil dan terbesar
	for i := 1; i < N; i++ {
		if weights[i] < minWeight {
			minWeight = weights[i]
		}
		if weights[i] > maxWeight {
			maxWeight = weights[i]
		}
	}

	// Output hasil
	fmt.Printf("Berat terkecil: %.2f\n", minWeight)
	fmt.Printf("Berat terbesar: %.2f\n", maxWeight)
}
