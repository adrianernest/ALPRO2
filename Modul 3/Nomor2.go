package main

import "fmt"

func f(x int) int { return x * x }
func g(x int) int { return x - 2 }
func h(x int) int { return x + 1 }

func main() {
	var a, b, c int
	if _, err := fmt.Scan(&a, &b, &c); err != nil {
		fmt.Println("Input error")
		return
	}

	fmt.Println(f(g(h(a))))
	fmt.Println(g(h(f(b))))
	fmt.Println(h(f(g(c))))
}
