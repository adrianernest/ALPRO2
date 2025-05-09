package main

import "fmt"

func urutkanArray(arr []int) {
	for i := 1; i < len(arr); i++ {
		kunci := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > kunci {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = kunci
	}
}

func periksaJarak(arr []int) (bool, int) {
	if len(arr) < 2 {
		return true, 0
	}
	jarak := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != jarak {
			return false, 0
		}
	}
	return true, jarak
}

func main() {
	var angka int
	var data []int

	for {
		fmt.Scan(&angka)
		if angka < 0 {
			break
		}
		data = append(data, angka)
	}

	urutkanArray(data)

	for i := 0; i < len(data); i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(data[i])
	}
	fmt.Println()

	berjarakTetap, jarak := periksaJarak(data)

	if berjarakTetap {
		fmt.Printf("Data berjarak %d\n", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
