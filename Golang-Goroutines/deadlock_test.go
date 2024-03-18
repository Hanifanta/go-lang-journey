package Golang_Goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	user1.Change(-amount)
	fmt.Println("Lock ", user1.Name)

	time.Sleep(1 * time.Second)

	user2.Lock()
	user2.Change(amount)
	fmt.Println("Lock ", user2.Name)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Hanif",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Al Irsyad",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)

	time.Sleep(3 * time.Second)

	fmt.Println("User 1 balance : ", user1.Balance)
	fmt.Println("User 2 balance : ", user2.Balance)
}
