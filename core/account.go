package   包 core

import (   导入(
    "crypto/ecdsa"
    "math/big"
    "web3sim/utils"
)

type Account struct {
    Address    string
    PrivateKey *ecdsa.PrivateKey
    Balance    *big.Int
}

func NewAccount() *Account {
    pk, addr := utils.GenerateKeyPair()
    return &Account{
        Address:    addr,
        PrivateKey: pk,
        Balance:    big.NewInt(1000000000000000000),
    }
}


