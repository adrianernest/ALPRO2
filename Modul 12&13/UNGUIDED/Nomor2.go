package main

import "fmt"

func selectionSortAsc(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		idxMin := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[idxMin] {
				idxMin = j
			}
		}
		arr[i], arr[idxMin] = arr[idxMin], arr[i]
	}
}

func selectionSortDesc(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		idxMax := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[idxMax] {
				idxMax = j
			}
		}
		arr[i], arr[idxMax] = arr[idxMax], arr[i]
	}
}

func bacaInput() [][]int {
	var jumlahDaerah, jumlahRumah int
	fmt.Scan(&jumlahDaerah)
	semuaDaerah := make([][]int, jumlahDaerah)

	for i := 0; i < jumlahDaerah; i++ {
		fmt.Scan(&jumlahRumah)
		ganjil := []int{}
		genap := []int{}

		for j := 0; j < jumlahRumah; j++ {
			var nomor int
			fmt.Scan(&nomor)
			if nomor%2 == 0 {
				genap = append(genap, nomor)
			} else {
				ganjil = append(ganjil, nomor)
			}
		}

		selectionSortAsc(ganjil)
		selectionSortDesc(genap)

		semuaDaerah[i] = append(ganjil, genap...)
	}

	return semuaDaerah
}

func tampilkanOutput(daerah [][]int) {
	for i := 0; i < len(daerah); i++ {
		for j := 0; j < len(daerah[i]); j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(daerah[i][j])
		}
		fmt.Println()
	}
}

func main() {
	data := bacaInput()
	fmt.Println()
	tampilkanOutput(data)
}
