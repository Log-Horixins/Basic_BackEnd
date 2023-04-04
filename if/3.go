package main

import "fmt"

func main() {
	var pulkam string

	fmt.Print("Apakah Ajia sering pulang kampung ? (y/n)")
	fmt.Scan(&pulkam)

	if pulkam == "y" {
		fmt.Println("Ajia kenna sanksi")
	} else {
		fmt.Println("Ajia orang baik")
	}
}
