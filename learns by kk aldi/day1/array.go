package main

import "fmt"

func main() {
	var nama [3]string
	nama[0] = "Coconut"
	nama[1] = "Computer"
	nama[2] = "Club"
	fmt.Println(nama)

	var months = [...]string{
		"january", "february", "maret", "april", "mei", "juni", "juli", "agustus", "september", "oktober", "november", "desember",
	}
	fmt.Println(months)
	fmt.Println("jumlah bulan : ", len(months))
}
