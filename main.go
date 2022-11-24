package main

import (
	"fmt"
	"go-second/library"
	. "go-second/library"
	"reflect"
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
	fmt.Println("jari-jari", (bangunDatar).keliling())
	fmt.Println("interface------------------------------")

	var rahasia interface{}

	rahasia = "ini rahasia"
	fmt.Println("isi dari ", rahasia)

	rahasia = []string{"data", "tipe data", "fungsi"}
	fmt.Println("nilai dari ", rahasia)
	fmt.Println("Any Interface------------------------------")
	var data map[string]interface{}

	data = map[string]interface{}{
		"buah":  "buah naga",
		"angka": 1929,
		"array": []string{"buah naga", "enak sekali"},
	}

	fmt.Println("interface", data)
	fmt.Println("Any Interface custom------------------------------")

	var number = "1929"
	var reflectType = reflect.TypeOf(number)
	var reflectValue = reflect.ValueOf(number)
	fmt.Println("tipe dari number ", number, " adalah", reflectType)
	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai dari number adalah ", reflectValue)
	} else {
		fmt.Println("nilai dari number adalah ", reflectValue)
	}
	fmt.Println("Reflect------------------------------")
	var nilaiNumber = "2919"
	var vNilai = reflect.ValueOf(nilaiNumber)
	fmt.Println("nilai", vNilai.Interface().(string))
	fmt.Println("tipe nilai", vNilai.Type())

	fmt.Println("Reflect x Interface------------------------------")
	var ugm = &pelajarMahasiswa{NamaKampus: "UGM", JumlahFakultas: 12}
	ugm.getPropertyInfo()
	fmt.Println("Reflect x Interface------------------------------")

}
