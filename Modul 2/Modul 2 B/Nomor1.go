package main

import "fmt"

func main() {
    referensi := []string{"merah", "kuning", "hijau", "ungu"}
    berhasil := true

    for i := 0; i < 5; i++ {
        for j := 0; j < 4; j++ {
            var warna string
            fmt.Scan(&warna)
            if warna != referensi[j] {
                berhasil = false
            }
        }
    }

    fmt.Println("BERHASIL:", berhasil)
}
