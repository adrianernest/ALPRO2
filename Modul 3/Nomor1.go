package main

import "fmt"

func factorial(n int) (f int) {
	f = 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	return
}

func perm(n, r int) int {
	if n < r {
		return 0
	}
	return factorial(n) / factorial(n-r)
}

func comb(n, r int) int {
	if n < r {
		return 0
	}
	return perm(n, r) / factorial(r)
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	if a >= c && b >= d {
		fmt.Println(perm(a, c), comb(a, c))
		fmt.Println(perm(b, d), comb(b, d))
	}
}
