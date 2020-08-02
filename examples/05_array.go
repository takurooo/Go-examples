package main

import "fmt"

func main() {
	{
		var arr [1]int
		fmt.Printf("%T %v %v %v\n", arr, arr, cap(arr), len(arr)) // [1]int [0]
	}
	{
		var arr [2]int = [2]int{1, 2}
		fmt.Printf("%T %v %v %v\n", arr, arr, cap(arr), len(arr)) // [2]int [1 2] 2 2
	}
	{
		arr := [...]string{"x", "y", "z"}
		fmt.Printf("%T %v %v %v\n", arr, arr, cap(arr), len(arr)) // [3]string [x y z] 3 3
	}
	{
		var arr [0]int
		fmt.Printf("%T %v %v %v\n", arr, arr, cap(arr), len(arr)) // [0]int [] 0 0
	}
	{
		arr := new([3]int) // pointer to an array
		arr[0] = 1
		arr[1] = 2
		arr[2] = 3
		fmt.Printf("%T %p %v %v %v\n", arr, arr, arr, cap(arr), len(arr)) // *[3]int 0xc0000141a0 &[1 2 3] 3 3
	}
}
