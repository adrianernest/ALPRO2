package main

import "fmt"

func main() {
	var a, b, c, d, e int
	var x, y, z byte

	fmt.Scan(&a, &b, &c, &d, &e)
	fmt.Scanln()

	fmt.Scanf("%c%c%c" , &x, &y, &z)

	fmt.Printf("%c%c%c%c%c\n", a, b, c, d, e)

	fmt.Printf("%c%c%c\n", x+1, y+1, z+1)
}
