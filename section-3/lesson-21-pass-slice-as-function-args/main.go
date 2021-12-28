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

	foo3 := []string{"a", "b"}
	fmt.Printf("foo3[0] address: %p \n", &foo3[0])

	// bar是一個pointer到foo3 array[0]的slice
	// 所以修改bar foo3也會被修改, 因為實際上是pointer到相同的array
	bar := foo3[:1]
	fmt.Println("bar:", bar)
	fmt.Printf("bar[0] address: %p \n", &bar[0])

	s1 := append(bar, "c")
	fmt.Println("foo3:", foo3)
	fmt.Printf("s1 address: %p \n", &s1)
	fmt.Printf("bar address: %p \n", &bar)

	fmt.Println("s1:", s1)
	fmt.Printf("s1[0] address: %p \n", &s1[0])

	s2 := append(bar, "d")
	fmt.Println("foo3:", foo3)
	fmt.Println("s2:", s2)

	// 超過foo3的capacity, append會回傳一個新的array
	s3 := append(bar, "e", "f")
	fmt.Println("foo3:", foo3)
	fmt.Println("s3:", s3)
	fmt.Printf("foo3[0] address: %p \n", &foo3[0])
	fmt.Printf("s3[0] address: %p \n", &s3[0])
}
