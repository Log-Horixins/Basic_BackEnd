package main

import "fmt"

func main() {
	var (
		uang  int
		ulang string
	)
RE:
	fmt.Println("Welcome to my Bank")

	fmt.Print("Masukkan jumlah uang (min Rp50000) : ")
	fmt.Scan(&uang)

	if uang >= 50000 {
		fmt.Println("Pemasukan uang berhasil")
		fmt.Println("Saldo pada akun anda sebesar Rp", uang)
	} else {
		fmt.Println("Uang yang anda masukkan kurang dari Nominal")
		fmt.Println("apakah anda ingin memasukkan ulang ? (y/n)")

		fmt.Scan(&ulang)

		if ulang == "y" {
			goto RE
		} else {
			goto end
		}
	}
end:
}
