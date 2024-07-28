package main

import "fmt"

func main() {
    // Deklarasi variabel string
    var nama string = "John Doe"
    fmt.Println(nama) // Output: John Doe

    // Deklarasi variabel string dengan cara shorthand
    namaLengkap := "John Doe"
    fmt.Println(namaLengkap) // Output: John Doe

    // Menggunakan tanda kutip ganda (") untuk string
    var kalimat string = "Halo, saya John Doe."
    fmt.Println(kalimat) // Output: Halo, saya John Doe.

    // Menggunakan tanda kutip tunggal (') untuk karakter
    var karakter rune = 'A'
    fmt.Println(karakter) // Output: 65 (kode ASCII dari 'A')

    // Menggunakan fungsi len() untuk menghitung panjang string
    var panjang int = len(nama)
    fmt.Println(panjang) // Output: 8

    // Menggunakan operator + untuk menggabungkan string
    var namaDepan string = "John"
    var namaBelakang string = "Doe"
    var namaLengkap2 string = namaDepan + " " + namaBelakang
    fmt.Println(namaLengkap2) // Output: John Doe

    // Menggunakan fungsi fmt.Sprintf() untuk menggabungkan string
    var namaLengkap3 string = fmt.Sprintf("%s %s", namaDepan, namaBelakang)
    fmt.Println(namaLengkap3) // Output: John Doe
}