package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan teks (akhiri dengan titik): ")
	input, _ := reader.ReadString('.')

	*n = 0 // Reset panjang array
	for _, char := range input {
		if unicode.IsLetter(char) { // Hanya masukkan huruf
			t[*n] = unicode.ToLower(char) // Ubah ke huruf kecil
			*n++
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(t[i]))
	}
	fmt.Println()
}

func balikkanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrome(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	isiArray(&tab, &m)

	fmt.Print("Teks yang dimasukkan: ")
	cetakArray(tab, m)

	balikkanArray(&tab, m)
	fmt.Print("Teks setelah dibalik: ")
	cetakArray(tab, m)

	if palindrome(tab, m) {
		fmt.Println("Palindrome? true")
	} else {
		fmt.Println("Palindrome? false")
	}
}
