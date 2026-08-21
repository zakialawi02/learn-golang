package main

import "fmt"

func main() {
	fmt.Println("Math Opration")

	var a = 10
	var b = 5
	var c = 2
	fmt.Println("Penjumlahan: ", a+b)
	fmt.Println("Pengurangan: ", a-b)
	fmt.Println("Perkalian: ", a*b)
	fmt.Println("Campuran: ", a-b+a*c)   // perkalian di dahulukan, mengikuti aturan matematika
	fmt.Println("Campuran: ", (a-b+a)*c) // operasi dalam kurung di dahulukan
	fmt.Println("Pembagian: ", a/b)
	fmt.Println("Modulo/sisa pembagian: ", a%b)

	// Augumented Assignment
	// adalah cara singkat untuk melakukan operasi matematika pada variabel
	fmt.Println("Augumented Assignment")
	a += 10 // sama dengan a = a + 10
	fmt.Println(a)
	b -= 5 // sama dengan b = b - 5
	fmt.Println(b)
	a *= 2 // sama dengan a = a * 2
	fmt.Println(a)
	b /= 10
	fmt.Println(b)
	a %= 3
	fmt.Println(a)

	// Unary Oprator
	fmt.Println("Unary Oprator")
	var x = 5
	var y = 2
	x++ // sama dengan x = x + 1
	fmt.Println(x)
	y-- // sama dengan y = y - 1
	fmt.Println(y)
}
