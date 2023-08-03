package main

import "fmt"

// FUnction Map

// len(map) 						= Untuk mendapatkan jumlah data di map
// map(key) 						= Mengambil data di map dengan key
// "len(map) = value" 				= Mengubah data di map dengan key
// make(map[TypeKey]TypeValue) 		= Membuat map baru
// delete(map, key) 				= Menghapus data di map dengan key

func main() {
	//-1
	person := map[string]string{
		"name":    "dayat",
		"address": "patowonua",
	}

	person["title"] = "student university of Muhammadiyah Makassar"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	//-2
	var book map[string]string = make(map[string]string)
	book["title"] = "Learn at Go-Lang"
	book["student"] = "Dayat"
	book["place"] = "Sekretariat"
	fmt.Println(book)
	fmt.Println(len(book))
	fmt.Println(book["title"])
	fmt.Println(book["student"])
	fmt.Println(book["place"])

	delete(book, "place")

	fmt.Println(book)
	fmt.Println(len(book))
	fmt.Println(book["title"])
	fmt.Println(book["student"])
	fmt.Println(book["place"])
}
