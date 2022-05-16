package main

import (
	"fmt"

	"github.com/DongjinS/learngo/part2/banking"
)

func main() {
	account := banking.Account{Owner: "DJ", Balance:100000000}

	fmt.Println(account, account.Owner)
}