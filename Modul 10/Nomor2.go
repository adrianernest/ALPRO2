package main

import (
	"fmt"
)

func main() {
	var x, y int

	// Input jumlah ikan (x) dan kapasitas wadah (y)
	fmt.Scan(&x, &y)

	// Input berat ikan
	weights := make([]float64, x)
	for i := 0; i < x; i++ {
		fmt.Scan(&weights[i])
	}

	// Menghitung total berat di setiap wadah
	var containers []float64
	var sum float64
	for i, weight := range weights {
		sum += weight
		if (i+1)%y == 0 || i == x-1 {
			containers = append(containers, sum)
			sum = 0
		}
	}

	// Output total berat di setiap wadah
	fmt.Println("Total berat di setiap wadah:")
	for i, container := range containers {
		fmt.Printf("Wadah %d: %.2f\n", i+1, container)
	}

	// Menghitung rata-rata berat di setiap wadah
	var total float64
	for _, container := range containers {
		total += container
	}
	average := total / float64(len(containers))

	// Output rata-rata berat
	fmt.Printf("Rata-rata berat di setiap wadah: %.2f\n", average)
}
