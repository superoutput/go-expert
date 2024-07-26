package main

import(
	"fmt"
)

type Pixel struct {
    X float64
    Y float64
}

type Vertex struct {
    X float64
    Y float64
    Z float64
}

func main() {
	/**
	How to initialize a array of interface in golang?
	references: https://stackoverflow.com/questions/52150146/how-to-initialize-a-array-of-interface-in-golang
	*/
	var i interface{}
	describe(i)

	i = []Pixel{
		{X: 12.5, Y: 23.4},
		{X: 17.2, Y: 7.9},
	}
	describe(i)

	i = []interface{}{ 
		Pixel{X: 12.5, Y: 23.4},
		Pixel{X: 17.2, Y: 7.9},
	}
	describe(i)

	i = []Vertex{
		{X: 10.7, Y: 13.3, Z: 25.1},
		{X: 18.3, Y: 16.9, Z: 16.4},
	}
	describe(i)

	i = []interface{}{
		Vertex{X: 10.7, Y: 13.3, Z: 25.1},
	}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	/**
	Why golang struct array cannot be assigned to an interface array
	reference: https://stackoverflow.com/questions/44319906/why-golang-struct-array-cannot-be-assigned-to-an-interface-array
	*/
	var x = []Vertex{
		{X: 10.7, Y: 13.3, Z: 25.1},
		{X: 18.3, Y: 16.9, Z: 16.4},
	}
	var y = make([]interface{}, len(x))
	for i, v := range x {
		y[i] = v
	}
	describe(x)
	describe(y)

	fmt.Println(len(x))
	fmt.Println(len(y))
	for _, item := range x {
		describe(item)
	}
	for _, item := range y {
		describe(item)
	}
	
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}