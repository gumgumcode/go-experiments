package main

import (
	"fmt"
	"sync"
	"time"
)

type BankAccount struct {
	balance int
	mu      sync.Mutex
}

func (acc *BankAccount) Deposit(amount int) {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	acc.balance += amount
	fmt.Printf("Deposited: %d. New balance: %d \n", amount, acc.balance)
}

func (acc *BankAccount) Withdraw(amount int) {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	if acc.balance >= amount {
		acc.balance -= amount
		fmt.Printf("Withdrew: %d. New balance: %d \n", amount, acc.balance)
	} else {
		fmt.Printf("Insufficient balance: %d \n", acc.balance)
	}
}

func main() {
	account := &BankAccount{balance: 2000}

	for i := 0; i < 5; i++ {
		go func() {
			account.Deposit(200)
		}()
		go func() {
			account.Withdraw(400)
		}()
	}

	// give time for go routines to complete
	time.Sleep(time.Second * 2)

	fmt.Printf("Final balance: %d", account.balance)
}
