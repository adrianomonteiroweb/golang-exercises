package main

import (
	"fmt"
)

type newType int

var x newType;
var y int;

func main() {
	y = int(x)

	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
