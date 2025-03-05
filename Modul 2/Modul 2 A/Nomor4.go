package main

import "fmt"

func main() {
    var celsius float64

    fmt.Print("Masukkan suhu dalam Celsius: ")
    fmt.Scan(&celsius)

    fahrenheit := (celsius * 9 / 5) + 32
    reamur := celsius * 4 / 5
    kelvin := celsius + 273.15

    fmt.Printf("Fahrenheit: %.2f°F\n", fahrenheit)
    fmt.Printf("Reamur: %.2f°R\n", reamur)
    fmt.Printf("Kelvin: %.2fK\n", kelvin)
}
