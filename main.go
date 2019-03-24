package main

import "fmt"

func main() {
	fmt.Println("hello")

	go func() {
		fmt.Println(Add(4.4, 90))
	}()

}

func Add(x, y float64) float64 {

	return x + y
}
