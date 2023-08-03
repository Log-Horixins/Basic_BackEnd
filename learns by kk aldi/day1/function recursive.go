package main

import "fmt"

func angkafaktorial(angka int) int {
	if angka == 1 {
		return 1
	} else {
		return angka * (angkafaktorial(angka - 1))
	}
}

func main() {
	var (
		// angka int
		total int
	)
	fmt.Print("masukkan angka: ")
	fmt.Scanln(&total)

	// for angka = total; angka > 1; angka-- {
	// 	fmt.Print(angka, "x")
	// 	total = total * (angka - 1)
	// }

	fmt.Println()
	total = 1

	fmt.Print(":", total)
}
