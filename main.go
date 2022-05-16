package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/DongjinS/learngo/something"
)

func multiply(a int,b int) int {
	return (a*b)
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	fmt.Println("Hello World")
	something.SayHello()
	fmt.Println(math.Phi)
	// var name string = "DJ"
	name := "DJ"
	name = "jja"
	fmt.Println(name)
	fmt.Println(multiply(2,2))
	totalLength, upperName := lenAndUpper("dongjin")
	fmt.Println(totalLength, upperName)
}