package main

import (
	"fmt"

	_ "github.com/scott/init-func/bar"
	_ "github.com/scott/init-func/foo"
)

var global = conver()

// init is the key work in golang
// init func執行時機是在main func之行之前才會執行
// 所以global一開始=100 -> init func global = 0 -> main func
func init() {
	fmt.Println("init global is ", global)
	global = 0
}

func conver() int {
	return 100
}

func main() {
	fmt.Println("run main")
	fmt.Println("main global is ", global)
}
