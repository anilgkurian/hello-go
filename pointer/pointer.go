package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

func ValueFunction(v Vertex) {
	v.X = 30 	
}

func PointerFunction(v *Vertex) {
	v.Y = 40 	
}

func (v Vertex) ValueMethod() {
	v.X = 50 	
}

func (v *Vertex) PointerMethod() {
	v.Y = 60 	
}


type MyType string

func (t MyType) add(x,y int) int{
	return x+y+2;
}

func main() {

	t1 := MyType("");
	fmt.Println(t1.add(1,2))

	v1 := Vertex{3, 4}
	ValueFunction(v1)
	PointerFunction(&v1)
	fmt.Println(v1)

	v2 := Vertex{5, 6}
	v2.ValueMethod()
	v2.PointerMethod()
	fmt.Println(v2)

}
