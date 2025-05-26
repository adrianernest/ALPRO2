package main

import (
	"fmt"
	"strings"
)

var kalimat string
var posisi int
var karakterSekarang byte

func mulai() {
	posisi = 0
	karakterSekarang = kalimat[posisi]
}

func maju() {
	posisi++
	if posisi < len(kalimat) {
		karakterSekarang = kalimat[posisi]
	}
}

func karakterSaatIni() byte {
	return karakterSekarang
}

func akhirKalimat() bool {
	return karakterSekarang == '.'
}

func main() {
	fmt.Print("Masukkan kalimat yang diakhiri dengan tanda titik ('.'): ")
	fmt.Scanln(&kalimat)

	if !strings.HasSuffix(kalimat, ".") {
		fmt.Println("Kalimat harus diakhiri dengan tanda titik ('.')")
		return
	}

	mulai()

	totalKarakter := 0
	totalA := 0
	totalLE := 0
	var karakterSebelumnya byte = 0

	for !akhirKalimat() {
		karakter := karakterSaatIni()
		totalKarakter++

		if strings.ToUpper(string(karakter)) == "A" {
			totalA++
		}

		if strings.ToUpper(string(karakterSebelumnya)) == "L" && strings.ToUpper(string(karakter)) == "E" {
			totalLE++
		}

		karakterSebelumnya = karakter
		maju()
	}

	frekuensiA := float64(totalA) / float64(totalKarakter) * 100

	fmt.Println("Total karakter    :", totalKarakter)
	fmt.Println("Total huruf 'A'   :", totalA)
	fmt.Printf("Frekuensi 'A'     : %.2f%%\n", frekuensiA)
	fmt.Println("Total kata 'LE'   :", totalLE)
}
