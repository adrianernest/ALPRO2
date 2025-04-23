package main

import "fmt"

func pencarianHuruf(kalimat string, karakter string) []int {
	var posisi []int

	for i := 0; i < len(kalimat); i++ {
		if kalimat[i] == karakter[0] {
			posisi = append(posisi, i)
		}
	}
	return posisi
}
func main() {

	var kalimat, karakter string

	fmt.Print("Kalimat: ")
	fmt.Scan(&kalimat)

	fmt.Print("Karakter: ")
	fmt.Scan(&karakter)

	posisi := pencarianHuruf(kalimat, karakter)

	if len(posisi) > 0 {
		fmt.Println("Posisi huruf", karakter, "dalam kalimat:", posisi)
	} else {
		fmt.Println("Karakter", karakter, "tidak ditemukan dalam kalimat.")
	}
}
