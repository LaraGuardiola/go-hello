package main // declares main package

import (
	"fmt" //imports formatter package
)

//*VARIABLES AND ARRAYS
var n int = 1
var notInfered = 2
var arr = []string{"a", "b", "c"}

//*CLASSES
type Pet struct {
	name string
	age  int
}

var instance = Pet{"dog", 2}

//* LOOPS, has the classic for, for-range(js forEach), while, do-while

//forEach loop like in js _ means we don't care about index, second is the elem of the array

func goLoop() {
	for _, e := range arr {
		fmt.Println(e)
	}
}

//everything has to be executed from main()

func main() {
	n = 3
	m := 2
	z := multiplier(n)
	//notInfered = "bruh" go has inferred type so it cannot be a string
	goLoop()
	fmt.Println("hello world", n+m+z)
	fmt.Scanf("%s", &instance.name)
	fmt.Println(arr)

}

func multiplier(x int) int {
	return x * 2
}
