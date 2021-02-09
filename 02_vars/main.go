package main

import "fmt"

func main()  {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128
	
	//	Using var
	// var name  = "ItsMurumba"
	var age  = 37
	const isCool = true

	// Shorthand
	// name := "ItsMurumba"
	size := 1.3
	// email := "itsmurumba@gmail.com"

	name, email := "ItsMurumba", "itsmurumba@gmail.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)
	
}