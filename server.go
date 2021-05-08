package main

import (
	"fmt"
	"net/http"
)

type boluint string

func louk(writer http.ResponseWriter, request *http.Request) {
	name := boluint("arewa")
	fmt.Fprintf(writer, "Hello World, %s!", string(name))
}
func main() {
	x := test()
	x.eat()
	x.walk(10)
	http.HandleFunc("/", louk)
	http.ListenAndServe(":8080", nil)
}

type animal interface {
	eat() int
	walk(uint) int
	sleep()
}

type elephant struct{}
type rat struct{}

func (e elephant) eat() int {
	return 500
}

func (e elephant) walk(d uint) int {
	return 500 * int(d)
}

func (r rat) eat() int {
	return 100
}

func (r rat) walk(d uint) int {
	return 100 * int(d)
}

func (r rat) boreHole() int {
	return 250
}

func test() animal {
	var pa rat
	// var ca kunle
	return pa
}
