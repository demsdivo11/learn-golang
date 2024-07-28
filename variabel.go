package main

import "fmt"


func main() {
	var nama string // var namaVariabel tipe data
	nama = "Dhemas" // deklarasi variabel nama : wajib tipe data yang sama
	fmt.Println(nama) // wajib memamggil variabel tsb

	nama = "Dhemas Dhiyanugraha"
	fmt.Println(nama)

	// cara lain agar tidak perlu memangil tipe data
	var dems = "Dhemas Divo"
	fmt.Println(dems)
}