package main

import "fmt"

func main() {
	// Comparison Operator adalah operator untuk membandingkan dua nilai
	// Operator ini mengembalikan nilai boolean (true atau false)
	// ada 6 jenis comparison operator: ==, !=, >, <, >=, <=

	// == : sama dengan (hasil: true jika nilai sama, false jika nilai beda)
	// != : tidak sama dengan (hasil: true jika nilai beda, false jika nilai sama)
	// >  : lebih besar dari (hasil: true jika nilai kiri lebih besar dari nilai kanan)
	// <  : lebih kecil dari (hasil: true jika nilai kiri lebih kecil dari nilai kanan)
	// >= : lebih besar dari atau sama dengan (hasil: true jika nilai kiri lebih besar dari atau sama dengan nilai kanan)
	// <= : lebih kecil dari atau sama dengan (hasil: true jika nilai kiri lebih kecil dari atau sama dengan nilai kanan)

	fmt.Println("COMPARISON OPERATOR")

	var a = 10
	var b = 5
	var c = 10
	fmt.Println(a == c)
	fmt.Println(a == b)
	fmt.Println(a != c)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a >= c)
	fmt.Println(a >= b)
	fmt.Println(a <= c)
	fmt.Println(a <= b)

	var d = "satu"
	var e = "dua"
	fmt.Println(d == e)
	fmt.Println(d != e)
	fmt.Println(d > e)
	fmt.Println(d < e)
	fmt.Println(d >= e)
	fmt.Println(d <= e)

	fmt.Println("BOOLEAN OPERATOR")
	// Boolean Operator adalah operator yang digunakan untuk operasi logika pada nilai boolean
	// Operator ini mengembalikan nilai boolean (true atau false)
	// ada 3 jenis boolean operator: &&, ||, !

	// && : AND (hasil: true jika kedua nilai true, false jika salah satu false)
	// || : OR  (hasil: true jika salah satu nilai true, false jika keduanya false)
	// !  : NOT (hasil: kebalikan dari nilai boolean)

	var nilaiAnd = true && true
	fmt.Println(nilaiAnd)
	var nilaiOr = true || false
	fmt.Println(nilaiOr)
	var nilaiNot = !true
	fmt.Println(nilaiNot)
}
