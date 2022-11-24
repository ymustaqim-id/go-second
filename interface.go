package main

import (
	"fmt"
	"math"
	"reflect"
)

type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}
func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

type pelajarMahasiswa struct {
	NamaKampus     string
	JumlahFakultas int
}

func (p *pelajarMahasiswa) getPropertyInfo() {
	var ambilNilai = reflect.ValueOf(p)

	if ambilNilai.Kind() == reflect.Ptr {
		ambilNilai = ambilNilai.Elem()
	}
	var ambilNilaiType = ambilNilai.Type()

	for i := 0; i < ambilNilaiType.NumField(); i++ {
		fmt.Println("nama", ambilNilai.Field(i))
		fmt.Println("type", ambilNilai.Field(i).Type())
		fmt.Println("nilai", ambilNilai.Field(i).Interface())
		fmt.Println("")
	}
}
