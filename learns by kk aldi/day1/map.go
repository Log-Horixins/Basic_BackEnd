package main

import "fmt"

func main() {

	var biodata = map[string]string{
		"nama":   "dayat",
		"status": "mahasiswa",
		"umur":   "19",
	}

	fmt.Println("nama: ", biodata["nama"])
	fmt.Println("status: ", biodata["status"])
	fmt.Println("umur: ", biodata["umur"])

}
