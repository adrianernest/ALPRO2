package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Domino struct {
	kiri  int
	kanan int
}

func buatSetDomino() []Domino {
	var setDomino []Domino
	for i := 0; i <= 6; i++ {
		for j := i; j <= 6; j++ {
			setDomino = append(setDomino, Domino{i, j})
		}
	}
	return setDomino
}

func kocokDomino(kartu []Domino) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(kartu), func(i, j int) {
		kartu[i], kartu[j] = kartu[j], kartu[i]
	})
}

func cocok(kartu Domino, ujung int) bool {
	return kartu.kiri == ujung || kartu.kanan == ujung
}

func balikKartu(kartu Domino) Domino {
	return Domino{kartu.kanan, kartu.kiri}
}

func tampilkanKartu(kartu Domino) string {
	return fmt.Sprintf("[%d|%d]", kartu.kiri, kartu.kanan)
}

func main() {
	setDomino := buatSetDomino()
	ronde := 0
	totalSkor := 0

	for {
		ronde++
		fmt.Printf("\n--- Ronde %d ---\n", ronde)
		kocokDomino(setDomino)

		tumpukan := append([]Domino(nil), setDomino...)
		kartuPemain := tumpukan[:7]
		tumpukan = tumpukan[7:]

		rantai := []Domino{tumpukan[0]}
		tumpukan = tumpukan[1:]
		skor := 0

		for {
			fmt.Println("\nRantai:")
			for _, kartu := range rantai {
				fmt.Print(tampilkanKartu(kartu), " ")
			}
			fmt.Println("\n\nKartu Anda:")
			for idx, kartu := range kartuPemain {
				fmt.Printf("%2d: %s\n", idx+1, tampilkanKartu(kartu))
			}
			fmt.Println("101010: Akhiri Ronde | 909090: Akhiri Permainan")
			fmt.Print("Pilihan Anda (-7..-1 kiri, 1..7 kanan): ")

			var pilihan int
			fmt.Scan(&pilihan)

			if pilihan == 101010 {
				fmt.Println("Ronde dihentikan.")
				break
			} else if pilihan == 909090 {
				fmt.Println("Permainan selesai.")
				fmt.Printf("Total Ronde: %d, Total Skor: %d\n", ronde-1, totalSkor)
				return
			}

			if pilihan < -len(kartuPemain) || pilihan > len(kartuPemain) || pilihan == 0 {
				fmt.Println("Pilihan tidak valid.")
				continue
			}

			absPilihan := pilihan
			if absPilihan < 0 {
				absPilihan = -absPilihan
			}
			if absPilihan > len(kartuPemain) {
				fmt.Println("Indeks kartu tidak tersedia.")
				continue
			}

			kartuTerpilih := kartuPemain[absPilihan-1]
			bisaMain := false

			if pilihan < 0 {
				ujungKiri := rantai[0].kiri
				if cocok(kartuTerpilih, ujungKiri) {
					if kartuTerpilih.kanan == ujungKiri {
						rantai = append([]Domino{kartuTerpilih}, rantai...)
					} else {
						rantai = append([]Domino{balikKartu(kartuTerpilih)}, rantai...)
					}
					bisaMain = true
				}
			} else {
				ujungKanan := rantai[len(rantai)-1].kanan
				if cocok(kartuTerpilih, ujungKanan) {
					if kartuTerpilih.kiri == ujungKanan {
						rantai = append(rantai, kartuTerpilih)
					} else {
						rantai = append(rantai, balikKartu(kartuTerpilih))
					}
					bisaMain = true
				}
			}

			if bisaMain {
				kartuPemain = append(kartuPemain[:absPilihan-1], kartuPemain[absPilihan:]...)
				skor++
			} else {
				if len(tumpukan) > 0 {
					fmt.Println("Kartu tidak cocok. Mengambil kartu dari tumpukan...")
					kartuPemain = append(kartuPemain, tumpukan[0])
					tumpukan = tumpukan[1:]
				} else {
					fmt.Println("Kartu tidak cocok dan tumpukan sudah habis.")
				}
			}

			if len(kartuPemain) == 0 {
				fmt.Println("Semua kartu habis. Ronde selesai!")
				break
			}

			bisaMainLagi := false
			ujungKiri := rantai[0].kiri
			ujungKanan := rantai[len(rantai)-1].kanan
			for _, kartu := range kartuPemain {
				if cocok(kartu, ujungKiri) || cocok(kartu, ujungKanan) {
					bisaMainLagi = true
					break
				}
			}
			if !bisaMainLagi && len(tumpukan) == 0 {
				fmt.Println("Tidak ada kartu yang bisa dimainkan dan tumpukan habis. Ronde selesai!")
				break
			}
		}

		fmt.Printf("Skor Ronde %d: %d\n", ronde, skor)
		totalSkor += skor
	}
}
