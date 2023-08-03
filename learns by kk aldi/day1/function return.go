package main

import "fmt"

func getsalam(nama string) string {
	return "hello " + nama
}

func getFullname(firstname string, lastname string) (string, string) {
	return firstname, lastname
}

func gettotal(nama string, angka ...int) (string, int) {
	var total int

	for i := 0; i < len(angka); i++ {
		total += angka[i]
	}

	return nama, total
}
func main() {
	salam := getsalam("dayat")
	fmt.Println(salam)

	firstname, lastname := getFullname("Nur", "Hidayat")
	namadepan, _ := getFullname("Nur", "Hidayat")

	fmt.Println("nama depan:", firstname)
	fmt.Println("nama belakang:", lastname)
	fmt.Println("nama depan:", namadepan)

	nama, total := gettotal("dayat", 7, 95, 3, 54, 223)
	fmt.Println("Nilai", nama, "adalah", total)
}
