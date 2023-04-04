package main

import "fmt"

func main() {
	// luas lingkaran
	const Pi float32 = 3.14
	var (
		r float32
	)

	fmt.Scan(&r)
	luas_lingkaran := Pi * r * r

	fmt.Println("Luas lingkaran adalah ", luas_lingkaran)

	// luas segitiga
	const per int64 = 1 / 2
	var (
		a int64
		t int64
	)

	fmt.Scan(&a)
	fmt.Scan(&t)
	luas_segitiga := per * a * t

	fmt.Println("Luas Segitiga adalah ", luas_segitiga)

	// luas kubus
	var s int64

	fmt.Scan(&s)
	luas_kubus := s * s

	fmt.Println("Luas Kubus adalah ", luas_kubus)
}
