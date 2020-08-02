package main

import (
	"fmt"
	"strconv"
)

func main() {
	{
		// 数値型の変換
		var i64 int64 = 1
		var f64 float64 = float64(i64)
		fmt.Printf(" 1: %T %v %f\n", f64, f64, f64) // float64 1 1.000000
	}

	{
		var f32 float32 = 1.9
		var ui32 uint32 = uint32(f32)
		fmt.Printf(" 2: %T %v %d\n", ui32, ui32, ui32) // uint32 1 1
	}

	{
		// ASCII codeから文字列へのの変換
		var s string = string(65)
		fmt.Printf(" 3: %T %v %s\n", s, s, s) // string A A
	}

	{
		// 文字列と数値の変換
		atoi, err := strconv.Atoi("33")               // string to int
		itoa := strconv.Itoa(-33)                     // int to string
		fmt.Printf(" 4: %T %d %v\n", atoi, atoi, err) // int -33 <nil>
		fmt.Printf(" 5: %T %s\n", itoa, itoa)         // string -33
	}

	// bitサイズとn進数を指定して文字列を数字に変換
	{
		b, err := strconv.ParseBool("true")
		f64, err := strconv.ParseFloat("3.14", 64)
		i64, err := strconv.ParseInt("9223372036854775807", 10, 32) // int64 max value -> int32 max value
		ui64, err := strconv.ParseUint("33", 10, 64)
		fmt.Printf(" 6: %T %v %v\n", b, b, err)       // bool true <nil>
		fmt.Printf(" 7: %T %f %v\n", f64, f64, err)   // float64 3.140000 <nil>
		fmt.Printf(" 8: %T %d %v\n", i64, i64, err)   // int64 2147483647 <nil>
		fmt.Printf(" 9: %T %d %v\n", ui64, ui64, err) // uint64 33 <nil>
	}
	{
		ui64, err := strconv.ParseUint("33a", 10, 64)
		fmt.Printf("10: %T %d %v\n", ui64, ui64, err) // int64 0 strconv.ParseUint: parsing "33a": invalid syntax
	}

	{
		s := strconv.Quote(`"Fran & Freddie's Diner	"`)
		fmt.Println(s)
	}

}
