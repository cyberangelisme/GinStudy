package main

import "fmt"

type sr interface {
	str2() string
}

type es struct {
	s string
}

func (e *es) str2() string {
	return e.s
}

func (e *es) Error2() string {
	return e.s
}

func test2(text string) sr {
	return &es{text}
}

func main() {
	fmt.Println(test2("aaa"))
	fmt.Println((&es{"bbb"}).str2())
}
