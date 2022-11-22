package main

import (
	"fmt"
	. "go-second/library"
)

func main() {
	SayHello("Ayu")
	var smk = Pelajar{Nama: "Ayu", Kelas: 12}
	fmt.Println("Nama", smk.Nama)
	fmt.Println("Kelas", smk.Kelas)
}
