package main

import "fmt"

func main() {
	{
		var m map[string]int // メモリは確保されない
		// panic: assignment to entry in nil map
		// m["x"] = 1
		fmt.Printf("%T %p %v\n", m, m, m) // map[string]int 0x0 map[]
	}
	{
		var m map[string]int = map[string]int{}
		m["x"] = 1
		m["y"] = 2
		m["z"] = 3
		for k, v := range m {
			fmt.Printf("%s %d\n", k, v)
		}
		{
			v := m["a"]
			fmt.Println(v) // 0
		}
		{
			// 値の存在をチェックして処理をする
			if v, ok := m["a"]; ok {
				fmt.Println(v, ok)
			} else {
				fmt.Println(v, ok) // 0 false
			}
		}
	}
	{
		var m map[string]int = map[string]int{"x": 1, "y": 2, "z": 3}
		fmt.Printf("%T %v\n", m, m) // map[string]int map[x:1 y:2 z:3]
	}
	{
		m := map[string]int{}
		m["xyz"] = 123
		m["abc"] = 123
		delete(m, "abc")
		fmt.Printf("%T %p %v\n", m, m, m) // map[string]int 0xc0000681e0 map[xyz:123]
	}
	{
		m := make(map[string]int)
		fmt.Printf("%T %p %v %v\n", m, m, m, len(m)) // map[string]int 0xc0000681e0 map[] 0
	}
	{
		m := make(map[string]int, 128)
		fmt.Printf("%T %p %v %v\n", m, m, m, len(m)) // map[string]int 0xc0000940f0 map[] 0
	}
}
