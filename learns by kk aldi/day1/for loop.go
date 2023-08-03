package main

import "fmt"

func main() {

	// nilai := 1
	// for nilai <= 5 {
	// 	fmt.Println(nilai)
	// 	nilai++
	// }

	// for nilai := 1; nilai <= 5; nilai++ {
	// 	fmt.Println(nilai)
	// }

	var months = [...]string{
		"january", "february", "maret", "april", "mei", "juni", "juli", "agustus", "september", "oktober", "november", "desember",
	}

	for i := 0; i < len(months); i++ {
		fmt.Println(months[i])
	}

	mhs := map[string]int{
		"misbah": 90,
		"hasrul": 89,
		"dayat":  10000,
		"ciwang": 1,
	}
	for key, values := range mhs {
		fmt.Println(key, ":", values)
	}
}
