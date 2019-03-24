package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello")

	go func() {
		fmt.Println(Add(4.4, 90))
	}()

	go func() {
		fmt.Println(Add(4.4, 90))
	}()

	go func() {
		fmt.Println(Add(4.4, 90))
	}()

	time.Sleep(time.Second * 10)
}

func Add(x, y float64) float64 {

	return x + y
}
