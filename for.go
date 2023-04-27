package main

import "fmt"

func main() {
	//individual declaration method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i +1
	}
	//for loop declaration in one line just like C.
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
	

