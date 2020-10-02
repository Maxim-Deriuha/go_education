package main

import "fmt"

func main() {
	var a int
	b := 0
	fmt.Println("введите а")
	fmt.Scan(&a)
	if a < b {
		fmt.Println(" а отрицательное ")
	} else if a == b {
		fmt.Println("a = 0")
	} else {
		fmt.Println("a положительное ")
	}

}
