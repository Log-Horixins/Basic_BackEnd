package main

import (
	"fmt"
)

//  nil = array, map, function

func bikinmap(kunci string, nilai string) map[string]string {
	if kunci == "" {
		return nil
	} else {
		return map[string]string{
			kunci: nilai,
		}
	}
}

func main() {
	biodata := bikinmap("nama", "hidayat")
	fmt.Println(biodata)
}
