package main

import "fmt"

func main() {
	nama := "hidayat"

	// switch length := len(nama); length > 10 {
	// case true:
	// 	fmt.Println("nama anda terlalu panjang")
	// case false:
	// 	fmt.Println("nama yang anda masukkan pas")
	// }

	length := len(nama)
	switch {
	case length > 10:
		fmt.Println("nama anda terlalu panjang")
	case length > 7:
		fmt.Println("nama anda kepanjangan")
	default:
		fmt.Println("nama anda memenuhi kondisi")
	}
}
