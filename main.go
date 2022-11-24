package main

import (
	"fmt"
	"go-second/library"
	. "go-second/library"
)

func main() {

	SayHello("Ayu")
	var smk = Pelajar{Nama: "Ayu", Kelas: 12}
	fmt.Println("Nama", smk.Nama)
	fmt.Println("Kelas", smk.Kelas)
	fmt.Println("exported unexported------------------------------")

	helloAyu("Name Ayu")
	fmt.Println("partial------------------------------")

	fmt.Println("Kampus", library.Mahasiswa.Kampus)
	fmt.Println("Jurusan", library.Mahasiswa.Jurusan)
	fmt.Println("init------------------------------")

	var bangunDatar hitung

	bangunDatar = persegi{12}
	fmt.Println("----------")
	fmt.Println("luas", bangunDatar.luas())
	fmt.Println("keliling", bangunDatar.keliling())
	fmt.Println("----------")

	bangunDatar = lingkaran{13}
	fmt.Println("luas", bangunDatar.luas())
	fmt.Println("keliling", bangunDatar.keliling())
	fmt.Println("----------")

	fmt.Println("luas", bangunDatar.luas())
	fmt.Println("keliling", bangunDatar.keliling())
	fmt.Println("jari-jari", (bangunDatar).keliling()
	fmt.Println("interface------------------------------")
}
