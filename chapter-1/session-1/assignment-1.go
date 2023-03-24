package main

import "fmt"

func main() {
	assignment()
}

func assignment() {
	var i int = 21
	var j bool = true
	var unicode rune = 'Я'
	var k float64 = 123.456

	// Menampilkan nilai i
	fmt.Printf("%v\n", i)

	// menampilkan tipe data dari variabel i
	fmt.Printf("%T\n", i)

	// menampilkan tanda %
	fmt.Printf("%%\n")

	//menampilkan nilai j : true
	fmt.Printf("%v\n", j)

	// menampilkan unicode russia : Я (ya)
	fmt.Printf("%c\n", unicode)
	fmt.Printf("%v\n", unicode)

	// menampilkan nilai base 10 : 21
	fmt.Printf("%d\n", i)

	// menampilkan nilai base 8 :25
	fmt.Printf("%o\n", i)

	//menampilkan nilai base 16 : f
	fmt.Printf("%x\n", i)

	//menampilkan nilai base 16 : F 13
	fmt.Printf("%X\n", i)

	//menampilkan unicode karakter Я : U+042F
	fmt.Printf("%U\n", unicode)

	//menampilkan float : 123.456000t
	fmt.Printf("%f\n", k)

	//menampilkan float scientific : 1.234560E+02
	fmt.Printf("%e\n", k)

}
