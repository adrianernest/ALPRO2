package main

import "fmt"

func sorting(arr []int) []int {
	ganjil := []int{}
	genap := []int{}

	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			genap = append(genap, arr[i])
		} else {
			ganjil = append(ganjil, arr[i])
		}
	}

	for i := 0; i < len(ganjil)-1; i++ {
		for j := i + 1; j < len(ganjil); j++ {
			if ganjil[j] < ganjil[i] {
				ganjil[i], ganjil[j] = ganjil[j], ganjil[i]
			}
		}
	}

	for i := 0; i < len(genap)-1; i++ {
		for j := i + 1; j < len(genap); j++ {
			if genap[j] > genap[i] {
				genap[i], genap[j] = genap[j], genap[i]
			}
		}
	}
	return append(ganjil, genap...)
}

func main() {
	angka := []int{12, 7, 3, 2, 9, 6, 8, 1, 11, 4}
	angkaTerurut := sorting(angka)
	fmt.Println(angkaTerurut)
}
