package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen dalam array: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("Isi array:", arr)

	fmt.Print("Elemen dengan indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("Elemen dengan indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var x int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Elemen dengan indeks kelipatan ", x, ": ")
	for i := x; i < n; i += x {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var index int
	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&index)
	if index >= 0 && index < n {
		arr = append(arr[:index], arr[index+1:]...)
		fmt.Println("Array setelah penghapusan:", arr)
	} else {
		fmt.Println("Indeks tidak valid.")
	}

	sum := 0
	for _, val := range arr {
		sum += val
	}
	rata := float64(sum) / float64(len(arr))
	fmt.Println("Rata-rata elemen array:", rata)

	var variance float64
	for _, val := range arr {
		variance += (float64(val) - rata) * (float64(val) - rata)
	}
	stdDeviasi := variance / float64(len(arr))
	fmt.Println("Standar deviasi elemen array:", stdDeviasi)

	var target int
	fmt.Print("Masukkan angka untuk menghitung frekuensinya: ")
	fmt.Scan(&target)
	count := 0
	for _, val := range arr {
		if val == target {
			count++
		}
	}
	fmt.Println("Frekuensi angka", target, "dalam array adalah:", count)
}
