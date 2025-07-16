package   包 core

import (   导入(
    "math/big"   “数学/大”
    "time"
    "github.com/google/uuid"
)

type Transaction struct {
    ID        string
    From      string
    To        string
    Amount    *big.Int
    Timestamp int64
    Status    string
}

func NewTransaction(from, to string, amount *big.Int) *Transaction {
    return &Transaction{
        ID:        uuid.New().String(),
        From:      from,
        To:        to,
        Amount:    amount,
        Timestamp: time.Now().Unix(),
        Status:    "pending",
    }
}


