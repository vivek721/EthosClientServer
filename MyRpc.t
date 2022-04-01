MyRpc interface { 
	getBalance(account Account) (uint64)
	transfer(acc1 Account, acc2 Account, amt uint64) (uint64, uint64)
}

Account struct {
	AccountID uint64
	AccountBalance uint64
}
