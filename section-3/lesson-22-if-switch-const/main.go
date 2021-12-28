package main

import "fmt"

func checkValue(s int) {
	switch s {
	case 2:
		fallthrough
	case 0, 1:
		fmt.Println("check value is ", s)
	}
}

func main() {
	checkValue(2)
	checkValue(1)
	checkValue(0)
}
