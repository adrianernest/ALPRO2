package main

import "fmt"

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

func daftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Scan(&(*pustaka)[i].id, &(*pustaka)[i].judul, &(*pustaka)[i].penulis, &(*pustaka)[i].penerbit, &(*pustaka)[i].eksemplar, &(*pustaka)[i].tahun, &(*pustaka)[i].rating)
	}
}

func cetakTerfavorit(pustaka DaftarBuku, n int) {
	terfavorit := pustaka[0]

	for i := 1; i < n; i++ {
		if pustaka[i].rating > terfavorit.rating {
			terfavorit = pustaka[i]
		}
	}

	fmt.Println("Buku terfavorit:")
	fmt.Printf("%s, %s, %s, %d, %d\n", terfavorit.judul, terfavorit.penulis, terfavorit.penerbit, terfavorit.tahun, terfavorit.rating)
}

func urutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		temp := (*pustaka)[i]
		j := i - 1
		for j >= 0 && (*pustaka)[j].rating < temp.rating {
			(*pustaka)[j+1] = (*pustaka)[j]
			j--
		}
		(*pustaka)[j+1] = temp
	}
}

func cetak5Terbaru(pustaka DaftarBuku, n int) {
	if n > 5 {
		n = 5
	}

	fmt.Println("\n5 Judul Buku dengan rating tertinggi:")

	for i := 0; i < n; i++ {
		fmt.Printf("%d. %s\n", i+1, pustaka[i].judul)
	}
}

func cariBuku(pustaka DaftarBuku, n, r int) {
	kiri, kanan := 0, n-1
	ditemukan := false
	var foundBooks []Buku

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if pustaka[tengah].rating == r {
			foundBooks = append(foundBooks, pustaka[tengah])
			i := tengah - 1
			for i >= 0 && pustaka[i].rating == r {
				foundBooks = append(foundBooks, pustaka[i])
				i--
			}
			i = tengah + 1
			for i < n && pustaka[i].rating == r {
				foundBooks = append(foundBooks, pustaka[i])
				i++
			}
			ditemukan = true
			break
		} else if pustaka[tengah].rating < r {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	if ditemukan {
		fmt.Printf("\nBuku dengan rating %d ditemukan:\n", r)
		for i := 0; i < len(foundBooks); i++ {
			fmt.Printf("%s, %s, %s, %d, %d, %d\n", foundBooks[i].judul, foundBooks[i].penulis, foundBooks[i].penerbit, foundBooks[i].tahun, foundBooks[i].eksemplar, foundBooks[i].rating)
		}
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var pustaka DaftarBuku
	var nPustaka, ratingCari int

	daftarkanBuku(&pustaka, &nPustaka)
	fmt.Scan(&ratingCari)
	urutBuku(&pustaka, nPustaka)
	cetakTerfavorit(pustaka, nPustaka)
	cetak5Terbaru(pustaka, nPustaka)
	cariBuku(pustaka, nPustaka, ratingCari)
}
