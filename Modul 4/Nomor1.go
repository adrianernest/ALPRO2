package main

import "fmt"

func factorial(n int) (hasil int) {
	hasil = 1
	for i := 2; i <= n; i++ {
		hasil *= i
	}
	return
}

func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func combination(n, r int) int {
	return permutation(n, r) / factorial(r)
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	fmt.Println(permutation(a, c), combination(a, c))
	fmt.Println(permutation(b, d), combination(b, d))
}
