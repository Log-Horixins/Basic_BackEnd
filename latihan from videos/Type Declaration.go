package main

import "fmt"

func main() {
	type huruf string
	type angka int
	type benar bool

	var nama huruf = "dayat"
	var nomor angka = 1234

	fmt.Println(nama)
	fmt.Println(nomor)
}
