package main

import "fmt"

func selectionSort(arr []int) {
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

func bacaInput() [][]int {
	var n, m int
	fmt.Scan(&n)
	semuaDaerah := make([][]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&m)
		rumah := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}
		selectionSort(rumah)
		semuaDaerah[i] = rumah
	}

	return semuaDaerah
}

func tampilkanOutput(semuaDaerah [][]int) {
	for i := 0; i < len(semuaDaerah); i++ {
		for j := 0; j < len(semuaDaerah[i]); j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(semuaDaerah[i][j])
		}
		fmt.Println()
	}
}

func main() {
	daerah := bacaInput()
	println()
	tampilkanOutput(daerah)
}
