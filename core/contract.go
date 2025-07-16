package   包 core

import (   导入(
	"web3sim/db"
)

type ContractEngine struct {
	State *db.StateDB
}

func (ce *ContractEngine) ExecuteTransfer(tx *Transaction) bool {
	return ce.State.Transfer(tx.From, tx.To, tx.Amount)
}


