package main

import (
	"fmt"
	"time"
)

type car struct {
	name string
}

// pass by value
// 會複製一個新的struct
func (c car) SetName01(s string) {
	fmt.Printf("1. address: %p \n", &c)
	c.name = s
}

// pass by address
func (c *car) SetName02(s string) {
	fmt.Printf("2. address: %p \n", c)
	c.name = s
}

type email struct {
	from string
	to   string
}

func (e *email) From(s string) {
	e.from = s
}

func (e *email) To(s string) {
	e.to = s
}

func (e *email) Send() {
	fmt.Printf("From: %s to: %s \n", e.from, e.to)
}

type email2 struct {
	from string
	to   string
}

func (e email2) From(s string) email2 {
	e.from = s
	return e
}

func (e email2) To(s string) email2 {
	e.to = s
	return e
}

func (e email2) Send() {
	fmt.Printf("From: %s to: %s \n", e.from, e.to)
}

func main() {
	c := &car{}
	fmt.Printf("c. address: %p \n", c)

	c.SetName01("foo")
	fmt.Println(c.name)

	c.SetName02("bar")
	fmt.Println(c.name)

	// goroutine 使用pass by address時要注意race condition的問題, 資料可能被別的協程改掉
	e := &email{}
	for i := 0; i < 10; i++ {
		go func(i int) {
			e.From(fmt.Sprintf("example%02d@com", i))
			e.To(fmt.Sprintf("example%02d@com", i+1))
			e.Send()
		}(i)
	}
	time.Sleep(1 * time.Second)

	// 使用pass by value的話可確保資料正確, struct method要記得加return
	e2 := &email2{}
	for i := 0; i < 10; i++ {
		go func(i int) {
			e2.From(fmt.Sprintf("example%02d@com", i)).
				To(fmt.Sprintf("example%02d@com", i+1)).
				Send()
		}(i)
	}
	time.Sleep(1 * time.Second)
}
