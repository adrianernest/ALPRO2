package main

import "fmt"

func main() {
	var clubA, clubB string
	fmt.Print("Klub A: ")
	fmt.Scan(&clubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&clubB)

	var winners []string
	matchNum := 1

	for {
		var scoreA, scoreB int
		fmt.Printf("Pertandingan %d: ", matchNum)
		fmt.Scan(&scoreA, &scoreB)

		if scoreA < 0 || scoreB < 0 {
			break
		}

		if scoreA > scoreB {
			winners = append(winners, clubA)
		} else if scoreA < scoreB {
			winners = append(winners, clubB)
		} else {
			winners = append(winners, "Draw")
		}

		matchNum++
	}

	fmt.Println("\nHasil pertandingan:")
	for i, winner := range winners {
		fmt.Printf("Hasil %d: %s\n", i+1, winner)
	}

	fmt.Println("Pertandingan selesai")
}
