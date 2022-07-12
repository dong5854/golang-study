package main

import (
	"fmt"
	"github.com/dong5854/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("dong")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		//log.Fatalln(err)
		fmt.Println(err)
	}
	account.ChangeOwner("lee")
	fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)
}
