package main

import "fmt"


func main()  {
	type huruf string
	var huruf1 huruf = "A"
	fmt.PrintLn(huruf1)

	type angka int
	var angka1 angka = 10
	fmt.Println(angka1)

	type trueOrFalse bool
	var benar trueOrFalse = false
	fmt.Println(trueOrFalse)
}