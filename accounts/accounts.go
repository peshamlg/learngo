package accounts

import (
  "errors"
  "fmt"
)

// Account struct
type Account struct {
  owner   string
  balance int
}

// NewAccount creates new account
func NewAccount(owner string) *Account {
  account := Account{owner: owner, balance: 0}
  return &account
}

// Owner of the account
func (a Account) Owner() string {
  return a.owner
}

// Balance of the account
func (a Account) Balance() int {
  return a.balance
}

func (a Account) String() string {
  return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) error {
  if (a.owner == newOwner) {
    return errors.New("New owner is same as previous")
  }
  a.owner = newOwner
  return nil
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) error {
  if amount > 0 {
     a.balance += amount
     return nil
  }
  reurn errors.New("Can't deposit less than 0")
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
  if a.balance < amount {
    return errors.New("Can't withdraw more than your balance")
  }
  a.balance -= amount
  return nil
}
