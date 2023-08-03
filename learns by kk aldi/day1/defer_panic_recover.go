package main

import "fmt"

func log() {
	pesanError := recover()
	if pesanError != nil {
		fmt.Println(pesanError)
	}
	fmt.Println("Aplikasi selesai ...")
}

func runApp(nilai int) {
	defer log()
	fmt.Println("Aplikasi berjalan ...")
	hasil := 1 / nilai
	if hasil != 0 {
		panic("")
	}
	fmt.Println("Hasil :", hasil)
}

func runApp2(error bool) {
	defer log()
	if error {
		panic("Aplikasi Error dengan inputan True")
	}
	fmt.Println("Aplikasi Berjalan ...")
}

func main() {
	runApp(0)
}
