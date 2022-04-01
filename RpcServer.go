package main

import (
	"ethos/altEthos"
	"ethos/syscall"
	"log"
	// "math/rand"
)

// var ACCOUNTS_LIST_SERVER []Account

func init() {

	// SetupMyRpcmakeAccount(makeAccount)
	SetupMyRpcgetBalance(getBalance)
	SetupMyRpctransfer(transfer)

}

func getBalance(account Account) MyRpcProcedure {
	log.Printf("MyRpcService : getBalance called\n")

	// // Get balance from account O(n)
	// for i := 0; i < len(ACCOUNTS_LIST); i++ {
	// 	if ACCOUNTS_LIST[i].AccountID == accountID {
	// 		return &MyRpcGetBalanceReply{ACCOUNTS_LIST[i].AccountBalance}
	// 	}
	// }
	// return &MyRpcGetBalanceReply{nil}

	return &MyRpcgetBalanceReply{account.AccountBalance}
}

// func makeAccount() MyRpcProcedure {

// 	initBalance := uint64(500)

// 	var account Account

// 	account.AccountID = 29375092375
// 	account.AccountBalance = initBalance

// 	// for i := 0; i < int(count); i++ {

// 	// 	// Generate a random balance for the account

// 	// 	// Append account to array
// 	// 	ACCOUNTS_LIST_SERVER = append(ACCOUNTS_LIST, account)
// 	// }

// 	return &MyRpcmakeAccountReply{account}
// }

func transfer(fromAcc Account, toAcc Account, amount uint64) MyRpcProcedure {
	if fromAcc.AccountBalance >= amount {
		fromAcc.AccountBalance -= amount
		toAcc.AccountBalance += amount
		log.Printf("Amount transferred between %v and %v\n", fromAcc.AccountID, toAcc.AccountID)
	} else {
		log.Printf("Could not transfer, negative\n")
	}

	return &MyRpctransferReply{fromAcc.AccountBalance, toAcc.AccountBalance}
}

func main() {

	altEthos.LogToDirectory("MyRpc/Server")

	listeningFd, status := altEthos.Advertise("MyRpc")
	if status != syscall.StatusOk {
		log.Printf("Advertising_service_failed:_%s\n", status)
		altEthos.Exit(status)
	}

	for {
		_, fd, status := altEthos.Import(listeningFd)
		if status != syscall.StatusOk {
			log.Printf("Error_calling_Import:_%v\n", status)
			altEthos.Exit(status)
		}
		log.Printf("listening fd %v", listeningFd)
		log.Printf("MyRpcService:_new_connection_accepted\n")

		t := MyRpc{}
		altEthos.Handle(fd, &t)
	}

}
