package main

import "fmt"

func main() {
	var a float32
	var h float32
	var l float32
	var h1 int
	var l1 int
	fmt.Print("введите число")
	fmt.Scan(&a)
	h = a / 30.0
	l = (a - h*30) * 0.5
	h1 = int(h)
	l1 = int(l)
	fmt.Print("It is ")
	fmt.Print(h1)
	fmt.Print("hours  ")
	fmt.Print(l1)
	fmt.Print("minutes.")
}
