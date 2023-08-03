package main

import "fmt"

func main() {
	var (
		p, l, m int
	)

	// fmt.Print("jumlah kertas: ")
	// fmt.Scanln(&jkertas)

	fmt.Print("panjang kertas: ")
	fmt.Scanln(&p)

	fmt.Print("lebar kertas: ")
	fmt.Scanln(&l)

	fmt.Print("jumlah lipatan: ")
	fmt.Scanln(&m)

	if (p%2 == 0) || (l%2 == 0) {
		if (p%2 == 0) && (l%2 == 0) {
			p /= m
			l /= m
		} else if p%2 == 0 {
			p /= m
			l = (l / m) - 1
		} else {
			p = (p / m) - 1
			l /= m
		}
	} else {
		p = (p / m) - 1
		l = (l / m) - 1
	}

	fmt.Println("panjang kertas ", p)
	fmt.Println("lebar kertas ", l)
	fmt.Println("jumlah pelipatan ", m)

}
