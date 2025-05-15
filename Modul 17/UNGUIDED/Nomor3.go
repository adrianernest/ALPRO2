package main

import (
	"fmt"
	"math/rand"
	"time"
)

func inisialisasi() int {
	var jumlahTetesan int
	fmt.Print("Masukkan jumlah tetesan air hujan: ")
	fmt.Scan(&jumlahTetesan)
	return jumlahTetesan
}

func tetesHujan(jumlahTetesan int) (int, int, int, int) {
	var daerahA, daerahB, daerahC, daerahD int
	for i := 0; i < jumlahTetesan; i++ {
		x := rand.Float64()
		y := rand.Float64()

		if x < 0.5 && y < 0.5 {
			daerahA++
		} else if x >= 0.5 && y < 0.5 {
			daerahB++
		} else if x < 0.5 && y >= 0.5 {
			daerahC++
		} else {
			daerahD++
		}
	}
	return daerahA, daerahB, daerahC, daerahD
}

func hitungCurahHujan(daerahA, daerahB, daerahC, daerahD int, ukuranTetesan float64) (float64, float64, float64, float64) {
	curahHujanA := float64(daerahA) * ukuranTetesan
	curahHujanB := float64(daerahB) * ukuranTetesan
	curahHujanC := float64(daerahC) * ukuranTetesan
	curahHujanD := float64(daerahD) * ukuranTetesan
	return curahHujanA, curahHujanB, curahHujanC, curahHujanD
}

func tampilkanHasil(curahHujanA, curahHujanB, curahHujanC, curahHujanD float64) {
	fmt.Printf("Curah hujan daerah A: %.3f milimeter\n", curahHujanA)
	fmt.Printf("Curah hujan daerah B: %.3f milimeter\n", curahHujanB)
	fmt.Printf("Curah hujan daerah C: %.3f milimeter\n", curahHujanC)
	fmt.Printf("Curah hujan daerah D: %.3f milimeter\n", curahHujanD)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ukuranTetesan := 0.0001

	jumlahTetesan := inisialisasi()
	daerahA, daerahB, daerahC, daerahD := tetesHujan(jumlahTetesan)

	curahHujanA, curahHujanB, curahHujanC, curahHujanD := hitungCurahHujan(daerahA, daerahB, daerahC, daerahD, ukuranTetesan)
	tampilkanHasil(curahHujanA, curahHujanB, curahHujanC, curahHujanD)
}
