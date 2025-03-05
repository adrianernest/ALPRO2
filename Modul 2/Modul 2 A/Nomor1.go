package main

import "fmt"

func main() {
    var (
        satu, dua, tiga string
        temp            string
    )

    fmt.Print("Masukkan input string: ")
    fmt.Scanln(&satu)
    fmt.Print("Masukkan input string: ")
    fmt.Scanln(&dua)
    fmt.Print("Masukkan input string: ")
    fmt.Scanln(&tiga)

    fmt.Println("Output awal =", satu, dua, tiga)

    temp = satu
    satu = dua
    dua = tiga
    tiga = temp

    fmt.Println("Output akhir =", satu, dua, tiga)
}

