package main

import "fmt"

func main() {
	var months = [...]string{
		"january", "february", "maret", "april", "mei", "juni", "juli", "agustus", "september", "oktober", "november", "desember",
	}
	fmt.Println(months)

	slice1 := months[3:6]
	fmt.Println(slice1)
	slice2 := months[3:]
	fmt.Println(slice2)
	slice3 := months[:10]
	fmt.Println(slice3)
	slice4 := months[:]
	fmt.Println(slice4)
}
