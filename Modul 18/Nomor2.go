package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Ubin struct {
	sisiSatu    int
	sisiDua     int
	nilai       int
	apakahGanda bool
}

type SetUbin struct {
	ubin   [28]Ubin
	jumlah int
}

func inisialisasiUbin(su *SetUbin) {
	idx := 0
	for a := 0; a <= 6; a++ {
		for b := a; b <= 6; b++ {
			u := Ubin{
				sisiSatu:    a,
				sisiDua:     b,
				nilai:       a + b,
				apakahGanda: a == b,
			}
			su.ubin[idx] = u
			idx++
		}
	}
	su.jumlah = 28
}

func acakUbin(su *SetUbin) {
	rand.Seed(time.Now().UnixNano())
	for i := su.jumlah - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		su.ubin[i], su.ubin[j] = su.ubin[j], su.ubin[i]
	}
}

func ambilUbin(su *SetUbin) Ubin {
	if su.jumlah == 0 {
		return Ubin{-1, -1, -1, false}
	}
	su.jumlah--
	return su.ubin[su.jumlah]
}

func dapatkanSisiUbin(u Ubin, sisi int) int {
	if sisi == 1 {
		return u.sisiSatu
	} else if sisi == 2 {
		return u.sisiDua
	} else {
		return -1
	}
}

func dapatkanNilaiUbin(u Ubin) int {
	return u.nilai
}

func cocok(u1, u2 Ubin) bool {
	return u1.sisiSatu == u2.sisiSatu ||
		u1.sisiSatu == u2.sisiDua ||
		u1.sisiDua == u2.sisiSatu ||
		u1.sisiDua == u2.sisiDua
}

func cariUbinCocok(su *SetUbin, referensi Ubin) Ubin {
	fmt.Println("\nUbin yang diambil: ")
	for su.jumlah > 0 {
		ubin := ambilUbin(su)
		fmt.Printf("Diambil: (%d,%d)\n", ubin.sisiSatu, ubin.sisiDua)
		if cocok(ubin, referensi) {
			return ubin
		}
	}
	return Ubin{-1, -1, -1, false}
}

func main() {
	var setUbin SetUbin

	inisialisasiUbin(&setUbin)
	acakUbin(&setUbin)

	ubinReferensi := ambilUbin(&setUbin)
	fmt.Printf("Ubin referensi: (%d,%d)\n", ubinReferensi.sisiSatu, ubinReferensi.sisiDua)

	ubinCocok := cariUbinCocok(&setUbin, ubinReferensi)

	if ubinCocok.sisiSatu == -1 {
		fmt.Println("Tidak ada ubin yang cocok.")
	} else {
		fmt.Printf("\nUbin yang cocok: (%d,%d)\n", ubinCocok.sisiSatu, ubinCocok.sisiDua)
	}
}
