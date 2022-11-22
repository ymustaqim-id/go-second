package library

import "fmt"

func SayHello(name string) {
	fmt.Print("Hello")
	introduceSelf(name)
}

func introduceSelf(name string) {
	fmt.Println("My name is", name)
}

type Pelajar struct {
	Nama  string
	Kelas int
}
