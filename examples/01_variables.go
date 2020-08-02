package main

import "fmt"

func main() {

	var (
		ui   uint   = 1                    // either 32 or 64 bits
		ui8  uint8  = 255                  // the set of all unsigned  8-bit integers (0 to 255)
		ui16 uint16 = 65535                // the set of all unsigned 16-bit integers (0 to 65535)
		ui32 uint32 = 4294967295           // the set of all unsigned 32-bit integers (0 to 4294967295)
		ui64 uint64 = 18446744073709551615 // the set of all unsigned 64-bit integers (0 to 18446744073709551615)

		uiptr uintptr = 1 // an unsigned integer large enough to store the uninterpreted bits of a pointer value

		i   int   = 1                    // either 32 or 64 bits
		i8  int8  = -128                 // the set of all signed  8-bit integers (-128 to 127)
		i16 int16 = -32768               // the set of all signed 16-bit integers (-32768 to 32767)
		i32 int32 = -2147483648          // the set of all signed 32-bit integers (-2147483648 to 2147483647)
		i64 int64 = -9223372036854775808 // the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

		f32 float32 = 1.1 // the set of all IEEE-754 32-bit floating-point numbers
		f64 float64 = 1.1 // the set of all IEEE-754 64-bit floating-point numbers

		c64  complex64  = complex(1, -1) // the set of all complex numbers with float32 real and imaginary parts
		c128 complex128 = complex(1, +1) // the set of all complex numbers with float64 real and imaginary parts

		b8 byte = 1 // alias for uint8
		r  rune = 1 // alias for int32

		s string = "abc"

		b bool = true
	)

	fmt.Printf("%T \t%v\n", ui, ui)
	fmt.Printf("%T \t%v\n", ui8, ui8)
	fmt.Printf("%T \t%v\n", ui16, ui16)
	fmt.Printf("%T \t%v\n", ui32, ui32)
	fmt.Printf("%T \t%v\n", ui64, ui64)
	fmt.Printf("%T \t%v\n", i, i)
	fmt.Printf("%T \t%v\n", i8, i8)
	fmt.Printf("%T \t%v\n", i16, i16)
	fmt.Printf("%T \t%v\n", i32, i32)
	fmt.Printf("%T \t%v\n", i64, i64)
	fmt.Printf("%T \t%v\n", f32, f32)
	fmt.Printf("%T \t%v\n", f64, f64)
	fmt.Printf("%T \t%v\n", uiptr, uiptr)
	fmt.Printf("%T \t%v\n", c64, c64)
	fmt.Printf("%T \t%v\n", c128, c128)
	fmt.Printf("%T \t%v\n", b8, b8)
	fmt.Printf("%T \t%v\n", r, r)
	fmt.Printf("%T \t%v\n", s, s)
	fmt.Printf("%T \t%v\n", b, b)

	// 関数内でのみ使用可能な省略系
	xi := 1
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Printf("%T %T %T %T %T\n", xi, xf64, xs, xt, xf) // int float64 string bool bool

}
