package main

import (
	"fmt"
	"os"
	"strconv"
)

type Player struct {
	Nama, Alamat, Pekerjaan, Alasan string
}

func main() {
	// untuk inputan argumen di CLI
	cliArguments := os.Args[1]

	// convert tipe data string menjadi integer
	noAbsen, _ := strconv.Atoi(cliArguments)

	// Data pemain bola
	var biodataPlayer = []Player{
		{
			Nama:      "Manuel Neuer",
			Alamat:    "Munchen",
			Pekerjaan: "Goal Keeper",
			Alasan:    "Mau jadi backend developer",
		},
		{
			Nama:      "Sven Ulreich",
			Alamat:    "Berlin",
			Pekerjaan: "Goal Keeper",
			Alasan:    "Mau jadi software engineer",
		},
		{
			Nama:      "Lucas Hernandez",
			Alamat:    "Paris",
			Pekerjaan: "Centre Back",
			Alasan:    "Ingin belajar buat server dan data base",
		},
		{
			Nama:      "Alphonso Davies",
			Alamat:    "Hamburg",
			Pekerjaan: "Left Back",
			Alasan:    "Ingin belajar tentang scalable web service",
		},
		{
			Nama:      "Benjamin Pavard",
			Alamat:    "Paris",
			Pekerjaan: "Right Back",
			Alasan:    "Ingin switch career jadi software engineer",
		},
		{
			Nama:      "Joshua Kimmich",
			Alamat:    "Munchen",
			Pekerjaan: "Defensive Midfielder",
			Alasan:    "Iseng ingin belajar Go-Lang",
		},
	}

	if noAbsen > len(biodataPlayer) {
		warningAlert()
	} else {
		showBiodataPlayer(biodataPlayer[noAbsen-1])
	}

}

func showBiodataPlayer(player Player) {
	fmt.Println("Nama :", player.Nama)
	fmt.Println("Alamat : ", player.Alamat)
	fmt.Println("Pekerjaan : ", player.Pekerjaan)
	fmt.Println("Alasan Memilih Go-Lang : ", player.Alasan)
}

func warningAlert() {
	fmt.Println("Tidak ada biodata player dalam no absen yang anda input")
}
