package main

import "fmt"

func main() {
	type NoKTP string
	// type [nama_alias] [tipe_data]
	// biasanya dgunakan untuk membuat alias dari tipe data yg sudah ada agar tidak terjadi kesalahan

	var ktpSaya NoKTP = "012304123324212"
	fmt.Println(ktpSaya)

	type Name string
	var nama Name = "Zaki"
	fmt.Println(nama)

	type NewName string
	// membuat tipe baru dari string, "string" bukan tipe yang sama dengan "NewName"
	var oldNama string = "zaki2"
	// updateNama adalah tipe NewName makanya harus dikonversi dulu
	var updateNama NewName = NewName(oldNama)
	fmt.Println(oldNama)
	fmt.Println(updateNama)

	type Umur int
	var umur Umur = 24
	fmt.Println(umur)
}
