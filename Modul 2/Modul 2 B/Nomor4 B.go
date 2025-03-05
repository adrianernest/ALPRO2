package main

import "fmt"

func main() {
	var K float64
	fmt.Print("Nilai K = ")
	fmt.Scan(&K)

	hasil := 1.0
	for i := 0.0; i < K; i++ {
		hasil *= ((4*i + 2) * (4*i + 2)) / ((4*i + 1) * (4*i + 3))
	}

	fmt.Printf("Nilai akar 2 = %.10f\n", hasil)
}
