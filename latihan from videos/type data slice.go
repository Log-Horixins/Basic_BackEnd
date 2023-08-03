package main

import "fmt"

// Function Slice

// len(slice) = untuk mendapatkan panjang
// cap(slice) = untuk mendapat kapasitas
// append(slice, data) = membuat slice baru dengan menambah data ke posisi terakhir slice, jika kapasitas sudah penuh, maka akan membuat array baru
// make([]TypeData, length, capacity) = membuat slice baru
// copy(destination, source) = menyalin slice dari source ke destination

func main() {
	var hari = [...]string{
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jum'at",
		"sabtu",
		"ahad",
	}

	var sclice1 = hari[2:5]
	fmt.Println(sclice1)
	fmt.Println(len(sclice1))
	fmt.Println(cap(sclice1))

	hari[3] = "diubah"
	fmt.Println(sclice1)

	sclice1[0] = "coconut"
	fmt.Println(hari)

	var slice2 = hari[5:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "eko")
	fmt.Println(slice3)
	slice3[1] = "bukan hari"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(hari)

	// Make Slice
	newSLice := make([]string, 2, 5)

	newSLice[0] = "ccnt"
	newSLice[1] = "coconut"

	fmt.Println(newSLice[1], newSLice[0], newSLice)
	fmt.Println(len(newSLice))
	fmt.Println(cap(newSLice))

	// copy slice ke slice lain
	copySlice := make([]string, len(newSLice), cap(newSLice))
	copy(copySlice, newSLice)
	fmt.Println(copySlice)

	// Perbedaan penulisan Array dan Slice
	// iniArray := [...]/[5]int{1,2,3,4,5}
	// iniSlice := []int{1,2,3,4,5}
}
