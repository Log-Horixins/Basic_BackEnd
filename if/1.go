package main

import "fmt"

func main() {
	var nilai uint8

	fmt.Scan(&nilai)

	if nilai <= 255 {
		fmt.Println("tipe data uint8")
	} else {
		fmt.Println("i don't know")
	}
}
