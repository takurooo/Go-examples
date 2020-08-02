package main

import "fmt"

func main() {

	// ------------------------------------
	// 型なし定数
	// ------------------------------------
	const userID = 100 // 型のない定数

	// userIDは型がないのでuint32にもuint64にも代入できる
	var ui32 uint32 = userID
	var ui64 uint64 = userID

	fmt.Println(ui32, ui64) // 100 100

	// ------------------------------------
	// 型付き定数
	// ------------------------------------
	const age uint = 10

	var ui uint = age

	fmt.Println(ui) // 10

	// ------------------------------------
	// Iota
	// ------------------------------------
	const (
		Sunday = iota // int
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays
	)

	// 0 1 2 3 4 5 6 7
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, numberOfDays)

}
