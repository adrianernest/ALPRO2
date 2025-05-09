package main

import "fmt"

func selectionSort(arr []int) {
	var i, j, minIdx, temp int
	for i = 0; i < len(arr)-1; i++ {
		minIdx = i
		for j = i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		temp = arr[i]
		arr[i] = arr[minIdx]
		arr[minIdx] = temp
	}
}

func hitungMedian(data []int) int {
	var n, tengah, median int
	n = len(data)
	selectionSort(data)

	if n%2 == 1 {
		tengah = n / 2
		median = data[tengah]
	} else {
		tengah = n / 2
		median = (data[tengah-1] + data[tengah]) / 2
	}
	return median
}

func main() {
	var input int
	var data []int
	var salinan []int
	var hasil int

	for {
		fmt.Scan(&input)

		if input < 0 {
			break
		} else if input == 0 {
			if len(data) > 0 {
				salinan = make([]int, len(data))
				copy(salinan, data)
				hasil = hitungMedian(data)
				fmt.Println(hasil)
			}
		} else {
			data = append(data, input)
		}
	}
}
