package main

import "fmt"

type Mahasiswa struct {
	nama string
	nim  string
}

func binarySearch(mahasiswa []Mahasiswa, nimCari string) int {
	kiri := 0
	kanan := len(mahasiswa) - 1

	for kiri <= kanan {
		mid := (kiri + kanan) / 2

		if mahasiswa[mid].nim == nimCari {
			return mid
		} else if mahasiswa[mid].nim < nimCari {
			kiri = mid + 1
		} else {
			kanan = mid - 1
		}
	}

	return -1
}

func main() {
	var X string

	mahasiswa := []Mahasiswa{
		{nama: "Faiq", nim: "103112430000"},
		{nama: "Adrian", nim: "103112430001"},
		{nama: "Rahmat", nim: "103112430002"},
		{nama: "Vincent", nim: "103112430003"},
	}
	X = "103112430001"

	fmt.Println("Data mahasiswa:", mahasiswa)

	index := binarySearch(mahasiswa, X)

	if index != -1 {
		fmt.Printf("Data ditemukan di array %d\n", index)
	} else {
		fmt.Println("Mahasiswa dengan NIM tersebut tidak ditemukan.")
	}
}
