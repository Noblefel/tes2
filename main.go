package main

import (
	"fmt"
	"strconv"
)

func main() {
	bersihkanLayar()

	for {
		fmt.Print("masukan jam kerja: ")

		jam, err := input()
		if err != nil || jam < 0 {
			bersihkanLayar()
			fmt.Println("kesalahan input, masukan angka yg benar!")
			continue
		}

		fmt.Print("masukan tarif: ")

		tarif, err := input()
		if err != nil || tarif < 0 {
			bersihkanLayar()
			fmt.Println("kesalahan input, masukan angka yg benar!")
			continue
		}

		gaji, gajiLembur := hitungGaji(jam, tarif)

		fmt.Println("=========== Hasil ============")
		fmt.Printf("Jam kerja : %.0f \n", jam)
		fmt.Printf("Tarif     : %.0f \n", tarif)
		fmt.Println("------------------------------")
		fmt.Printf("Gaji : Rp. %.0f \n", gaji)
		fmt.Printf("Gaji lembur (%.0f jam) : Rp. %.0f \n", max(jam-40, 0), gajiLembur)

		fmt.Println("...")
		input()
	}
}

func hitungGaji(jam, tarif float64) (float64, float64) {
	var gaji float64
	var gajiLembur float64

	if jam > 40 {
		gaji = 40 * tarif
		gajiLembur = (jam - 40) * tarif * 1.5
	} else {
		gaji = jam * tarif
	}

	return gaji, gajiLembur
}

func bersihkanLayar() {
	fmt.Print("\033[H\033[J")
}

func input() (float64, error) {
	var input string
	fmt.Scanln(&input)

	n, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	return float64(n), nil
}
