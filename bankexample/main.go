package main

import(
	"fmt"
	"github.com/peshamlg/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("walter")
	fmt.Println(account)

	errDeposit := account.Deposit(10)
	if errDeposit != nil {
		fmt.Println(errDeposit)
	}
	fmt.Println(account)

	errWithdraw := account.Withdraw(20)
	if errWithdraw != nil {
		fmt.Println(errWithdraw)
	}
	fmt.Println(account)
}
