package main

import "fmt"

// For Statement

// init statement	= statement sebelum for dieksekusi
// post statement	= statement yang akan selalu dieksekusi di akhir tiap perulangan

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter += 1
	}
	// atau
	for counter1 := 2; counter1 <= 16; counter1++ {
		fmt.Println("mengulang ke", counter1)
	}

	// for slice
	inislice := []string{"jamal", "kris", "akbar"}

	for i := 0; i < len(inislice); i++ {
		fmt.Println(inislice[i])
	}

	// for range
	for i, value := range inislice {
		fmt.Println("index", i, "=", value)
		// fmt.Println(value)
	}

	// for map
	person := make(map[string]string)
	person["name"] = "eko"
	person["title"] = "programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
