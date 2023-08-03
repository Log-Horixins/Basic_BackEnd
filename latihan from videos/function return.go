package main

import "fmt"

func gethello(nama string) string {
	if nama == "" {
		return "welcome to the jungle bro"
	} else {
		return "welcome to the jungle " + nama
	}
}

func main() {
	result := gethello("cione")

	fmt.Println(result)
	fmt.Println(gethello("paang"))
	fmt.Println(gethello(""))

}
