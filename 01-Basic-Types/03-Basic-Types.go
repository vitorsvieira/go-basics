package main

import "fmt"

func main() {

	//*************************************
	//bool declaration and initialization
	//*************************************
	var winter bool
	winter = false
	var spring bool = true
	var summer = true
	autumn := false

	fmt.Println("bool: ", winter, spring, summer, autumn)

	//*************************************
	//string declaration and initialization
	//*************************************
	var name string
	name = "Gopher"
	var lang string = "Golang"
	var food = "pasta"
	editor := "Sublime Text"

	fmt.Println("string: ", name, lang, food, editor)

	//*************************************
	//int declaration and initialization
	//same size as uint
	//*************************************
	var a, b, c int = 0, 2, 4
	var x, y, z = 8, 16, 32
	size := 64

	fmt.Println("int: ", a, b, c, x, y, z, size)

	//*************************************
	//int8 declaration and initialization
	//the set of all signed  8-bit integers (-128 to 127)
	//*************************************
	var posInt8 int8 = 30
	var negInt8 int8 = -50

	fmt.Println("int8: ", posInt8, negInt8)

	//*************************************
	//int16 declaration and initialization
	//the set of all signed 16-bit integers (-32768 to 32767)
	//*************************************
	var posInt16 int16 = 31500
	var negInt16 int16 = -24800

	fmt.Println("int16: ", posInt16, negInt16)

	//*************************************
	//int32 declaration and initialization
	//the set of all signed 32-bit integers (-2147483648 to 2147483647)
	//*************************************
	var posInt32 int32 = 2000000000
	var negInt32 int32 = -2000000000

	fmt.Println("int32: ", posInt32, negInt32)

	//*************************************
	//int64 declaration and initialization
	//the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
	//*************************************
	var posInt64 int64 = 7666555444333222111
	var negInt64 int64 = -7666555444333222111

	fmt.Println("int64: ", posInt64, negInt64)

	//*************************************
	//uint declaration and initialization
	//either 32 or 64 bits
	//*************************************
	var posUint uint = 7666555444333222111
	var zeroUint uint = 0

	fmt.Println("uint: ", posUint, zeroUint)

	//*************************************
	//uint8 declaration and initialization
	//the set of all unsigned  8-bit integers (0 to 255)
	//*************************************
	var posUint8 uint8 = 255
	var zeroUint8 uint8 = 0

	fmt.Println("uint8: ", posUint8, zeroUint8)

	//*************************************
	//uint16 declaration and initialization
	//the set of all unsigned 16-bit integers (0 to 65535)
	//*************************************
	var posUint16 uint16 = 55555
	var zeroUint16 uint16 = 0

	fmt.Println("uint16: ", posUint16, zeroUint16)

	//*************************************
	//uint32 declaration and initialization
	//the set of all unsigned 32-bit integers (0 to 4294967295)
	//*************************************
	var posUint32 uint32 = 3222111000
	var zeroUint32 uint32 = 0

	fmt.Println("uint32: ", posUint32, zeroUint32)

	//*************************************
	//uint64 declaration and initialization
	//the set of all unsigned 64-bit integers (0 to 18446744073709551615)
	//*************************************
	var posUint64 uint64 = 17666555444333222111
	var zeroUint64 uint64 = 0

	fmt.Println("uint64: ", posUint64, zeroUint64)

	//*************************************
	//uintptr declaration and initialization
	//an unsigned integer large enough to store the uninterpreted bits of a pointer value
	//*************************************
	var posUintptr uintptr = 7666555444333222111
	var oUintptr uintptr = 1e19

	fmt.Println("uintptr: ", posUintptr, oUintptr)

	//*************************************
	//byte(uint8) declaration and initialization
	//alias for uint8
	//*************************************
	var ch1 byte = 'a'
	var ch2 byte = 'A'
	var ch3 rune = 'δ'
	fmt.Println("byte: ", ch1, ch2, ch3) //print-> byte: 97 65 948

	//*************************************
	//rune(int8)(similar to char) declaration and initialization
	//alias for int32
	//*************************************
	var r1 rune = 'a'
	var r2 rune = 'δ'
	fmt.Println("rune: ", r1, r2) //print-> rune: 97 948

	//*************************************
	//float32 declaration and initialization
	//the set of all IEEE-754 32-bit floating-point numbers
	//(+- 1O-45 -> +- 3.4 * 1038 )
	//float32 is reliably accurate to about 7 decimal places
	//*************************************
	var posFloat32 float32 = 9.123456789  //4.123457
	var negFloat32 float32 = -9.123456789 //-4.123457

	fmt.Println("float32: ", posFloat32, negFloat32)

	//*************************************
	//float64 declaration and initialization
	//the set of all IEEE-754 64-bit floating-point numbers
	//(+- 5 * 10-324 -> 1.7 * 10308 )
	//float64 is reliably accurate to about 15 decimal places
	//Use float64 whenever possible, because all the functions of the math package expect that type
	//*************************************
	var posFloat64 float64 = 4.987654321123456789  //4.987654321123457
	var negFloat64 float64 = -4.987654321123456789 //-4.987654321123457

	fmt.Println("float64: ", posFloat64, negFloat64)

	//*************************************
	//complex64 declaration and initialization
	//the set of all complex numbers with float32 real and imaginary parts
	//*************************************
	var c64 complex64 = 5 + 10i
	fmt.Printf("complex64: %f\n", c64)

	//*************************************
	//complex128 declaration and initialization
	//the set of all complex numbers with float64 real and imaginary parts
	//*************************************
	var c128 complex128 = 10 + 100i + 1000i
	fmt.Printf("complex128: %f\n", c128)

}
