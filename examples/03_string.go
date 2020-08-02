package main

import (
	"fmt"
	"strings"
)

func main() {
	user := "abc"
	id := "123"
	userid := user + id
	fmt.Println(userid)            // abc123
	fmt.Println(userid[:])         // abc123(type string)
	fmt.Println(userid[0:1])       // a
	fmt.Println(userid[0])         // 97
	fmt.Println(string(userid[0])) // a

	// 大文字小文字変換
	fmt.Println(strings.ToUpper("toupper")) // TOUPPER
	fmt.Println(strings.ToUpper("TOLOWER")) // tolower

	// タイトルケースに変換
	fmt.Println(strings.Title("hey")) // "Hey"

	// 文字列を比較
	fmt.Println(strings.Compare("a", "a")) // 0
	fmt.Println(strings.Compare("a", "b")) // -1
	fmt.Println(strings.Compare("b", "a")) // 1
	fmt.Println("abc" == "abc")            // true

	// 指定した文字列が含まれているか
	fmt.Println(strings.Contains("abcdef\n", "\n")) // true

	// 指定した文字数をカウント
	fmt.Println(strings.Count("aaa", "a")) // 3

	// 空白文字でsplit
	fmt.Println(strings.Fields("ab cd ef")) // [ab cd ef] (type slice)

	// 先頭/終端文字列の一致確認
	fmt.Println(strings.HasPrefix("prefix_abc", "prefix")) // true
	fmt.Println(strings.HasPrefix("prefix_abc", "abc"))    // false
	fmt.Println(strings.HasSuffix("prefix_abc", "prefix")) // false
	fmt.Println(strings.HasSuffix("prefix_abc", "abc"))    // true

	// 文字列のインデックスを取得
	fmt.Println(strings.Index("abcc", "c"))     // 2
	fmt.Println(strings.Index("abcc", "z"))     // -1
	fmt.Println(strings.LastIndex("abcc", "c")) // 3

	// 文字列型のsliceを連結
	{
		s := []string{"a", "b", "c"}
		fmt.Println(strings.Join(s, ",")) // "a,b,c"
	}

	// 文字列置換
	fmt.Println(strings.Replace("abc abc", "a", "z", 1)) // "zbc abc"
	fmt.Println(strings.ReplaceAll("abc abc", "a", "z")) // "zbc zbc"

	// 文字列を分割
	fmt.Println(strings.Split("a,b,c", ","))      // [a b c]
	fmt.Println(strings.SplitAfter("a,b,c", ",")) // [a, b, c]

	// 文字列の削除
	fmt.Println(strings.Trim("aaxaa", "a"))        // "x"
	fmt.Println(strings.TrimLeft("aaxaa", "a"))    // "xaa"
	fmt.Println(strings.TrimRight("aaxaa", "a"))   // "aax"
	fmt.Println(strings.TrimRight("aaxaa", "xa"))  // ""
	fmt.Println(strings.TrimPrefix("aaxaa", "aa")) // "xaa"
	fmt.Println(strings.TrimSuffix("aaxaa", "aa")) // "aax"
	fmt.Println(strings.TrimSpace("  aaxaa  "))    // "aaxaa"

}
