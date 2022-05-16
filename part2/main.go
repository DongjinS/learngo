package main

import (
	"fmt"
	"log"

	"github.com/DongjinS/learngo/part2/accounts"
	"github.com/DongjinS/learngo/part2/mydict"
)

func main() {
	// bankingPrac()
	dictPrac()
}

func bankingPrac() {
	account := accounts.NewAccount("DJ")
	account.Deposit(10)
	fmt.Println(account)
	if err:=account.Withdraw(5); err!=nil{
		log.Fatalln(err)
	}
	account.ChangeOwner("HSE")
	fmt.Println(account.Owner())
	fmt.Println(account)
}

func dictPrac() {
	dictionary := mydict.Dictionary{"first": "First word"}
	
	definition, err := dictionary.Search("second")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	if err:=dictionary.Add("Hello", "Greeting");err!=nil{
		fmt.Println(err)
	}

	hello, err := dictionary.Search("Hello")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Found: Hello,", "Definition:", hello)
	}

	if err:=dictionary.Add("Hello", "Greeting");err!=nil{
		fmt.Println(err)
	}

	err = dictionary.Update("first", "not second")
	if err != nil {
		fmt.Println(err)
	}

	def,err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(def)
	}
	fmt.Println(dictionary)
	if err:=dictionary.Delete("hello"); err != nil {
		fmt.Println(err)
	}
	fmt.Println(dictionary)
}