package main

import (
	"ethos/altEthos"
	"ethos/syscall"
	"log"
)

var ACC Account

func init() {

	SetupMyRpcgetBalanceReply(getBalanceReply)
	SetupMyRpctransferReply(transferReply)
}

func getBalanceReply(accountBalance uint64) MyRpcProcedure {
	log.Printf("Client_Received_Balance_Reply with Balance: %v\n", accountBalance)
	return nil
}

func transferReply(amt1 uint64, amt2 uint64, status bool) MyRpcProcedure {
	if status == true {
		log.Printf("Transfer Successfull\n")
		log.Printf("New Amount for account 1 is: %v \n New Amount for account 2 is: %v\n", amt1, amt2)
	} else {
		log.Printf("Transfer failed\n")
	}
	return nil
}

func main() {
	altEthos.LogToDirectory("test/rpcClient")
	log.Printf("rpcClient:_before_call\n")

	acc1 := Account{AccountID: 1, AccountBalance: 100}
	acc2 := Account{AccountID: 2, AccountBalance: 50}
	fd1, status1 := altEthos.IpcRepeat("MyRpc", "", nil)
	if status1 != syscall.StatusOk {
		log.Printf("Ipc_failed:_%v\n", status1)
		altEthos.Exit(status1)
	}
	call1 := MyRpcgetBalance{acc1}
	status1 = altEthos.ClientCall(fd1, &call1)

	fd1, status1 = altEthos.IpcRepeat("MyRpc", "", nil)
	if status1 != syscall.StatusOk {
		log.Printf("Ipc_failed:_%v\n", status1)
		altEthos.Exit(status1)
	}

	call1 = MyRpcgetBalance{acc2}
	status1 = altEthos.ClientCall(fd1, &call1)
	if status1 != syscall.StatusOk {
		log.Printf("RpcClient: clientCall failed: %v\n", status1)
		altEthos.Exit(status1)
	}

	// TRANSFER
	fd2, status2 := altEthos.IpcRepeat("MyRpc", "", nil)
	if status1 != syscall.StatusOk {
		log.Printf("Ipc_failed:_%v\n", status2)
		altEthos.Exit(status2)
	}

	call2 := MyRpctransfer{acc1, acc2, uint64(10)}
	status2 = altEthos.ClientCall(fd2, &call2)

	if status2 != syscall.StatusOk {
		log.Printf("RpcClient: clientCall failed: %v\n", status2)
		altEthos.Exit(status2)
	}

	log.Printf("RpcClient:done\n")
}
