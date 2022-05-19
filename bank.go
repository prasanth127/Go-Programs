//Small Bank Application:

package main

import (
	"fmt"
	"sync"
	"time"
)

//Deposit
//Balance
//Saving Account Creation
//Withdrawal
var balance int
var wg sync.WaitGroup
var mut sync.Mutex
var rwmut sync.RWMutex

func Deposit(amount int) {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		mut.Lock()
		balance = balance + amount
		time.Sleep(time.Millisecond)
		mut.Unlock()
		wg.Done()
	}
}

func Withdrawal(amount int) bool {
	wg.Add(1)
	defer wg.Done()
	defer mut.Unlock()
	mut.Lock()
	balance = balance - amount
	if balance < 0 {
		balance = balance + amount
		///wg.Done()
		return false
	}

	time.Sleep(time.Millisecond)
	//wg.Done()

	return true
}

func Balance() int {
	rwmut.Lock()
	defer rwmut.Unlock()
	b := balance
	return b

}
func main() {

	//1.Check Balance
	fmt.Println("Balance = Rs.", Balance())
	//2.Deposit Rs.100
	go Deposit(100)
	fmt.Println("Deposited Rs.100")
	//3.Check Balance
	fmt.Println("Balance = Rs.", Balance())
	//4.Withdraw Rs.50
	for i := 0; i < 100; i++ {
		go Withdrawal(50)
		if balance > 0 {
			fmt.Println("Successful Withdrawal of Rs.50")
			fmt.Println("Balance = Rs.", Balance())
		} else {
			fmt.Println("Un-Successful Withdrawal of Rs.50")
			fmt.Println("Balance = Rs.", Balance())
		}
	}
	//5.Withdraw Rs.10
	if Withdrawal(10) {
		fmt.Println("Successful Withdrawal of Rs.10")
		fmt.Println("Balance = Rs.", Balance())
	} else {
		fmt.Println("Un-Successful Withdrawal of Rs.10")
		fmt.Println("Balance = Rs.", Balance())
	}
	//6.Withdraw Rs.100
	if Withdrawal(100) {
		fmt.Println("Successful Withdrawal of Rs.100")
		fmt.Println("Balance = Rs.", Balance())
	} else {
		fmt.Println("Un-Successful Withdrawal of Rs.100")
		fmt.Println("Balance = Rs.", Balance())
	}
	//7.Check the Balance
	fmt.Println("Balance = Rs.", Balance())
	//8.Deposit(Rs.100)
	Deposit(100)
	fmt.Println("Deposited Rs.100")
	fmt.Println("Balance =", Balance())
	//9.Deposit(Rs.200)
	Deposit(200)
	fmt.Println("Deposited Rs.200")
	fmt.Println("Balance = Rs.", Balance())

	wg.Wait()

}
