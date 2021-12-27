package main

import "fmt"

func modify(foo []string) {
	foo[1] = "c"
	fmt.Println("modify foo:", foo)
}

func addValue(foo []string) {
	foo = append(foo, "d")
	fmt.Println("modify foo:", foo, " len:", len(foo), "cap:", cap(foo))
}

func main() {
	foo := make([]string, 3)
	foo[0] = "a"
	foo[1] = "b"
	fmt.Println("len:", len(foo), "cap:", cap(foo))

	fmt.Println("before foo:", foo)
	modify(foo)
	fmt.Println("after foo:", foo)

	fmt.Println("before foo:", foo)
	addValue(foo)
	fmt.Println("after foo:", foo)
}
