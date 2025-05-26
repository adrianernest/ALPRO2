package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Domino struct {
	sisiSatu    int
	sisiDua     int
	nilai       int
	adalahBalak bool
}

type SetDomino struct {
	kartu  [28]Domino
	jumlah int
}

func inisialisasiSetDomino(set *SetDomino) {
	idx := 0
	for a := 0; a <= 6; a++ {
		for b := a; b <= 6; b++ {
			kartu := Domino{
				sisiSatu:    a,
				sisiDua:     b,
				nilai:       a + b,
				adalahBalak: a == b,
			}
			set.kartu[idx] = kartu
			idx++
		}
	}
	set.jumlah = 28
}

func acakKartu(set *SetDomino) {
	rand.Seed(time.Now().UnixNano())
	for i := set.jumlah - 1; i > 0; i-- {
		indeksAcak := rand.Intn(i + 1)
		set.kartu[i], set.kartu[indeksAcak] = set.kartu[indeksAcak], set.kartu[i]
	}
}

func ambilKartu(set *SetDomino) Domino {
	if set.jumlah == 0 {
		return Domino{-1, -1, -1, false}
	}
	set.jumlah--
	return set.kartu[set.jumlah]
}

func dapatkanSisiKartu(kartu Domino, sisi int) int {
	if sisi == 1 {
		return kartu.sisiSatu
	} else if sisi == 2 {
		return kartu.sisiDua
	}
	return -1
}

func dapatkanNilaiKartu(kartu Domino) int {
	return kartu.nilai
}

func main() {
	var setDomino SetDomino

	inisialisasiSetDomino(&setDomino)
	acakKartu(&setDomino)

	fmt.Println("Menampilkan 5 kartu pertama yang diambil:")
	for i := 0; i < 5; i++ {
		kartu := ambilKartu(&setDomino)
		fmt.Printf("Kartu %d: (%d,%d), Nilai: %d, AdalahBalak: %v\n",
			i+1, kartu.sisiSatu, kartu.sisiDua, dapatkanNilaiKartu(kartu), kartu.adalahBalak)
	}
}
