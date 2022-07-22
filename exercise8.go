package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for i2 := 1; i2 <=3; i2++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
