package main

import "fmt"

func main() {
	var (
		nilai1 int8
		nilai2 int8
	)

	nilai1 = 10
	nilai2 = 2

	hasil := nilai1 % nilai2

	fmt.Println("hasil dari nilai1 % nilai2 = ", hasil)
}
