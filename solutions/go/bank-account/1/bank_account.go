package account

import "sync"

// Define the Account type here.
type Account struct {
	amount int64
	lock   sync.Mutex
	closed bool
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{amount: amount}
}

func (a *Account) Balance() (int64, bool) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if a.closed {
		return 0, false
	}
	return a.amount, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	if a.closed {
		return 0, false
	}
	a.lock.Lock()
	defer a.lock.Unlock()
	nextAmount := a.amount + amount
	if nextAmount < 0 {
		return 0, false
	}
	a.amount = nextAmount
	return a.amount, true
}

func (a *Account) Close() (int64, bool) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if a.closed == true {
		return 0, false
	}
	a.closed = true
	amount := a.amount
	a.amount = 0
	return amount, true
}
