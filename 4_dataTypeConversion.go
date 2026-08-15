package main

import "fmt"

func main() {
	fmt.Println("DATA TYPE CONVERSION")

	// membuat variabel int32
	var nilai32 int32 = 321121
	fmt.Println(nilai32)

	// to int16 (max value int16: 32767)
	var nilai16 int16 = int16(nilai32)
	fmt.Println(nilai16) // melebihi batas max maka nilainya menjadi negatif

	// integer ke float
	var nilaiFloat32 float32 = float32(nilai32)
	fmt.Println(nilaiFloat32)

	// integer ke string
	var nilaiString string = string(nilai32)
	fmt.Println(nilaiString)

	// string functions, byte to string
	fmt.Println("STRING FUNC")
	fmt.Println("Hitung lenght string")
	var name = "HELLO ZAKI"
	fmt.Println(len(name))
	fmt.Println(name[3]) // byte karater index tersebut
	var getString = name[3]
	var string = string(getString)
	fmt.Println(getString)
	fmt.Println(string)
	fmt.Println(getString, "=", string)
}
