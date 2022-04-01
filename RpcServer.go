package main

import (
	"ethos/altEthos"
	"ethos/syscall"
	"log"
)

func init() {

	SetupMyRpcgetBalance(getBalance)
	SetupMyRpctransfer(transfer)

}

func getBalance(account Account) MyRpcProcedure {
	// log.Printf("MyRpcService : getBalance called\n")
	return &MyRpcgetBalanceReply{account.AccountBalance}
}

func transfer(fromAcc Account, toAcc Account, amount uint64) MyRpcProcedure {
	log.Printf("Balance for %v : %v\n", fromAcc.AccountID, fromAcc.AccountBalance)
	log.Printf("Balance for %v : %v\n", toAcc.AccountID, toAcc.AccountBalance)
	log.Printf("Amount of $%v transferred from %v to %v\n", amount, fromAcc.AccountID, toAcc.AccountID)
	var status bool
	if fromAcc.AccountBalance >= amount {
		fromAcc.AccountBalance -= amount
		toAcc.AccountBalance += amount
		status = true
	} else {
		log.Printf("Could not transfer, negative\n")
		status = false
	}

	return &MyRpctransferReply{fromAcc.AccountBalance, toAcc.AccountBalance, status}
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
		log.Printf("MyRpcService:_new_connection_accepted\n")

		t := MyRpc{}
		altEthos.Handle(fd, &t)
	}

}
