package main

import "fmt"

func main() {
	var awesome string

	fmt.Print("Apakah ciwang ganteng ? (y/n)")
	fmt.Scan(&awesome)

	if awesome == "y" {
		fmt.Println("Luar biasa")
	} else {
		fmt.Println("astaga")
	}
}
