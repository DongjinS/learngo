package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/DongjinS/learngo/part0-1/something"
)

func multiply(a int,b int) int {
	return (a*b)
}

// naked return - return 변수명 지정해서 리턴할 수 있음
func lenAndUpper(name string) (length int, uppercase string) {
	// defer ~~ call back
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	var res int
	for index, number := range numbers {
		fmt.Println(index, number)
		res+=number
	}
	for i:=0; i<len(numbers); i++{
		fmt.Println(numbers[i])
	}
	return res
}

func canIDrink(age int) bool {
	if koreanAge:=age+2; koreanAge<20{
		return false
	}
	return true
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
	fmt.Println(superAdd(1,2,3,4,5))
	fmt.Println(canIDrink(12))

	// pointer
	a:=2
	b:=&a
	fmt.Println(a,b,&a,&b,*b)

	// arrays
	names := [7]string{"nico", "lynn", "DJ", "SE", "HJ"}
	fmt.Println(names)
	// slices - array without length
	names2 := []string{"hi", "22"}
	names2 = append(names2, "123")
	fmt.Println(names2)

	// map

	nico := map[string]string {
		"name": "nico",
		"age": "12",
	}

	for key, value := range nico {
		fmt.Println(key, value)
	}

	// struct
	type person struct {
		name string;
		age int;
		favFood []string;
	}
	favFood := []string {"gogi","apple"}
	DJ := person{name: "DJ", age: 12, favFood: favFood}
	fmt.Println(DJ, DJ.favFood)

	// channels + go routines -> will learn from next projeccts
}