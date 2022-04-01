package main

import (
	"ethos/altEthos"
	"ethos/syscall"
	"log"
)

// type Account struct {
// 	AccountID      uint64
// 	AccountBalance uint64
// }

var ACC Account

func init() {

	SetupMyRpcgetBalanceReply(getBalanceReply)
	SetupMyRpctransferReply(transferReply)
}

func getBalanceReply(accountBalance uint64) MyRpcProcedure {
	log.Printf("Client_Received_Balance_Reply with ID:_%v\n", accountBalance)
	return nil
}

func transferReply(amt1 uint64, amt2 uint64) MyRpcProcedure {
	log.Printf("first one is: %v and second one is: %v\n", amt1, amt2)
	return nil
}

func main() {
	altEthos.LogToDirectory("test/rpcClient")
	log.Printf("rpcClient:_before_call\n")
	// fd, status := altEthos.IpcRepeat("MyRpc", "", nil)
	// if status != syscall.StatusOk {
	// 	log.Printf("Ipc_failed:_%v\n", status)
	// 	altEthos.Exit(status)
	// }
	// call := MyRpcmakeAccount{}
	// status = altEthos.ClientCall(fd, &call)

	acc1 := Account{AccountID: 1, AccountBalance: 100}
	acc2 := Account{AccountID: 2, AccountBalance: 50}
	fd1, status1 := altEthos.IpcRepeat("MyRpc", "", nil)
	if status1 != syscall.StatusOk {
		log.Printf("Ipc_failed:_%v\n", status1)
		altEthos.Exit(status1)
	}
	// log.Printf("hello fd : ", fd1, "\n")
	call1 := MyRpcgetBalance{acc1}
	status1 = altEthos.ClientCall(fd1, &call1)

	fd1, status1 = altEthos.IpcRepeat("MyRpc", "", nil)
	if status1 != syscall.StatusOk {
		log.Printf("Ipc_failed:_%v\n", status1)
		altEthos.Exit(status1)
	}
	// log.Printf("hello fd : ", fd1, "\n")
	call1 = MyRpcgetBalance{acc2}
	status1 = altEthos.ClientCall(fd1, &call1)

	// TRANSCER
	fd2, status2 := altEthos.IpcRepeat("MyRpc", "", nil)
	if status1 != syscall.StatusOk {
		log.Printf("Ipc_failed:_%v\n", status2)
		altEthos.Exit(status2)
	}
	// log.Printf("hello fd : ", fd1, "\n")
	call2 := MyRpctransfer{acc1, acc2, uint64(10)}
	status2 = altEthos.ClientCall(fd2, &call2)
	// status = altEthos.ClientCall(fd, &call)
	// status = altEthos.ClientCall(fd, &call)

	// status1 = altEthos.ClientCall(fd1, &call)
	// if status1 != syscall.StatusOk {
	// 	log.Printf("clientCall_failed:_%v\n", status1)
	// 	altEthos.Exit(status)
	// }

	log.Printf("RpcClient:done\n")
}
