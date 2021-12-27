package main

import "fmt"

func modify(foo []string) {
	foo[1] = "c"
	fmt.Println("modify foo:", foo)
}

func addValue(foo []string) {
	// 如果slice的cap不夠的話, append會回傳原本的array＊２
	// 新array的cap會是原本的兩倍
	foo = append(foo, "d")
	fmt.Println("modify foo:", foo)
}

func addValueByRef(foo *[]string) {
	*foo = append(*foo, "e")
}

func main() {
	foo := make([]string, 2, 4)
	foo[0] = "a"
	foo[1] = "b"
	fmt.Println("len:", len(foo))

	fmt.Println("before foo:", foo)
	modify(foo)
	fmt.Println("after foo:", foo)

	fmt.Println("before foo:", foo)
	addValue(foo)
	fmt.Println("after foo:", foo)

	foo2 := make([]string, 2, 3)
	foo2[0] = "a"
	foo2[1] = "b"

	// foo2 cap 還夠, 不會複製新的ARRAY
	fmt.Printf("%p \n", &foo2[0])
	addValueByRef(&foo2)
	fmt.Printf("%p \n", &foo2[0])

	// foo2 cap 滿了, Slice pointer到新的array
	fmt.Printf("%p \n", &foo2[0])
	addValueByRef(&foo2)
	fmt.Printf("%p \n", &foo2[0])
}
