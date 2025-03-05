package main

import "fmt"

func main() {
    var r float64
    const pi = 3.1415926535

    fmt.Print("Masukkan jari-jari bola: ")
    fmt.Scan(&r)

    r2 := r * r       
    r3 := r2 * r      

    volume := (4.0 / 3.0) * pi * r3
    luas := 4 * pi * r2

    fmt.Printf("Volume bola: %.2f\n", volume)
    fmt.Printf("Luas permukaan bola: %.2f\n", luas)
}
