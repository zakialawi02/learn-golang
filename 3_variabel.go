package main

import "fmt"

func main() {
	fmt.Println("Variabel")

	var nama string

	nama = "Zaki"
	fmt.Println(nama) // Zaki

	nama = "Zaki Alawi"
	fmt.Println(nama) // Zaki Alawi

	var umur int

	umur = 23
	fmt.Println("Umur: ", umur) // Umur:  23

	// tanpa declare type var dan langsung isi nilainya
	var nama_lengkap = "Zaki Alawi"
	fmt.Println(nama_lengkap)

	nama_lengkap = "Zaki Alawi EDIT"
	fmt.Println(nama_lengkap)

	// := adalah singkatan dari var dan := langsung meng-assign nilainya, hanya declare awal saja
	pekerjaan := "Programmer"
	fmt.Println(pekerjaan)

	pekerjaan = "Freelance"
	fmt.Println(pekerjaan)

	// declare multiple variable
	var (
		nama_depan    = "Zaki"
		nama_belakang = "Alawi"
		hobi          = "Baca Buku"
	)
	fmt.Println(nama_depan, nama_belakang)
	fmt.Println(hobi)

	// CONSTANT
	// const adalah variabel yang tidak bisa diubah nilainya
	const PI = 3.14159
	fmt.Println(PI)

	// multipe constant
	const (
		negara    = "Indonesia"
		mata_uang = "IDR"
		benua     = "Asia"
	)
	fmt.Println(negara, mata_uang, benua)
}
