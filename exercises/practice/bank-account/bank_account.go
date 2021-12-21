package account

import "sync"

// Define the Account type here.
type Account struct {
	total  int64
	closed bool
}

var accountLock sync.Mutex

func Open(amount int64) *Account {
	accountLock.Lock()
	defer accountLock.Unlock()
	if amount < 0 {
		return nil
	}
	return &Account{amount, false}
}

func (a *Account) Balance() (int64, bool) {
	accountLock.Lock()
	defer accountLock.Unlock()
	if a == nil || a.closed {
		return 0, false
	}
	return a.total, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	accountLock.Lock()
	defer accountLock.Unlock()
	if a == nil || a.closed {
		return 0, false
	}

	if amount < 0 && a.total+amount < 0 {
		return a.total, false
	}

	a.total = a.total + amount
	return a.total, true
}

func (a *Account) Close() (int64, bool) {
	accountLock.Lock()
	defer accountLock.Unlock()
	if !a.isActive() {
		return 0, false
	} else {
		a.closed = true
		return a.total, true
	}

}

func (a *Account) isActive() bool {
	// accountLock.Lock()
	// defer accountLock.Unlock()
	if a == nil || a.closed {
		return false
	} else {
		return true
	}
}
