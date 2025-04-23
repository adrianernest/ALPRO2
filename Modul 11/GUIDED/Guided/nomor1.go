package main

import "fmt"

type array []string

func dataBarang(n int) array {
	var barang string
	var arrayBarang array

	for i := 0; i < n; i++ {
		fmt.Print("Masukkan nama barang ke-", i+1, ": ")
		fmt.Scan(&barang)
		arrayBarang = append(arrayBarang, barang)
	}
	return arrayBarang
}

func cariArray(arrayBarang array, x string) bool {
	for i := 0; i < len(arrayBarang); i++ {
		if arrayBarang[i] == x {
			return true
		}
	}
	return false
}

func main() {
	var n int
	var x string

	fmt.Print("Masukkan jumlah barang : ")
	fmt.Scan(&n)

	data := dataBarang(n)
	fmt.Println("Data barang:", data)

	fmt.Print("Masukkan barang yang ingin dicari : ")
	fmt.Scan(&x)

	barang := cariArray(data, x)

	if barang {
		fmt.Println("Barang ditemukan!")
	} else {
		fmt.Println("Barang tidak ditemukan.")
	}
}
