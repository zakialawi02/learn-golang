package main

import "fmt"

func main() {
	fmt.Println("ARRAY DATA")
	// array adalah kumpulan data dengan tipe data yang sama
	// array memiliki ukuran yang tetap
	// [] adalah panjang/index array

	var nama [3]string // membuat array dengan tipe data string dan panjang index max 3, jika melebihi akan error
	// jika tidak diisi maka akan menjadi default value nya yaitu "" (string kosong)(string), atau 0 (integer) atau false (boolean) atau nil (pointer/slice/map/channel)

	nama[0] = "Zaki"
	nama[1] = "Alawi"
	nama[2] = "Budi"
	fmt.Println(nama)
	fmt.Println(nama[0])
	fmt.Println(nama[1])
	fmt.Println(nama[2])

	fmt.Println("Membuat array secara langsung")
	var values = [4]int{1, 2, 3, 4}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(values[3])
	// bisa gunakan [...] agar length otomatis mengikuti jumlah data yang diisi diawal. contoh: var nama = [...]string{"Zaki", "Alawi"} maka length nya adalah 2.

	fmt.Println("Fucntion Array")
	// len() adalah function untuk mendapatkan panjang array
	fmt.Println(len(nama))    // cek panjang array
	fmt.Println(len(nama[1])) // cek panjang index array
	nama[2] = "Andi"          // update value index array
	fmt.Println(nama)
}
