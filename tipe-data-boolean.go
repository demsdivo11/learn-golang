package main

import "fmt"

func main() {
    // Deklarasi variabel boolean
    var isAdmin bool = true
    
    // Cetak nilai isAdmin
    fmt.Println("Apakah admin?", isAdmin)
    
    // Contoh penggunaan boolean dalam kondisi if
    if isAdmin {
        fmt.Println("Anda adalah admin!")
    } else {
        fmt.Println("Anda bukan admin!")
    }
    
    // Contoh penggunaan boolean dalam operator logika
    var isMember bool = false
    if isAdmin && isMember {
        fmt.Println("Anda adalah admin dan member!")
    } else {
        fmt.Println("Anda tidak memenuhi kriteria!")
    }
    
    // Contoh penggunaan boolean dalam operator not
    var isNotAdmin bool = !isAdmin
    fmt.Println("Apakah bukan admin?", isNotAdmin)
}