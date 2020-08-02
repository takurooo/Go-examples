package main

import "fmt"

func main() {
	{
		var slice []int
		fmt.Printf("%T %v %v %v\n", slice, slice, cap(slice), len(slice)) // []int [] 0 0
	}
	{
		var slice []int = []int{1, 2}
		fmt.Printf("%T %v %v %v\n", slice, slice, cap(slice), len(slice)) // []int [1 2] 2 2
	}
	{
		slice := []string{}                                                         // not allocated yet
		fmt.Printf("%T %p %v %v %v\n", slice, slice, slice, cap(slice), len(slice)) // []string 0x1190fd0 [] 0 0
		slice = append(slice, "x")
		fmt.Printf("%T %p %v %v %v\n", slice, slice, slice, cap(slice), len(slice)) // []string 0xc000084030 [x] 1 1
		slice = append(slice, "y")
		fmt.Printf("%T %p %v %v %v\n", slice, slice, slice, cap(slice), len(slice)) // []string 0xc00008e0a0 [x y] 2 2
		slice = append(slice, "z")
		fmt.Printf("%T %p %v %v %v\n", slice, slice, slice, cap(slice), len(slice)) // []string 0xc00004e0c0 [x y z] 4 3
	}
	{
		slice := make([]int, 3, 4)                                                  // type len cap
		fmt.Printf("%T %p %v %v %v\n", slice, slice, slice, cap(slice), len(slice)) // []int 0xc000090000 [0 0 0] 4 3
	}
	{
		slice := make([]int, 2)                                                     // type len cap
		fmt.Printf("%T %p %v %v %v\n", slice, slice, slice, cap(slice), len(slice)) // []int 0xc000018080 [0 0] 2 2
	}
}
