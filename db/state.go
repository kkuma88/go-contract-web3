package   包 db

import (   导入(
    "math/big"   “数学/大”
    "sync"
)

type AccountState struct {
    Balance *big.Int
    Nonce   uint64
}

type StateDB struct {
    Accounts map[string]*AccountState
    mu       sync.RWMutex
}

func NewStateDB() *StateDB {
    return &StateDB{
        Accounts: make(map[string]*AccountState),
    }
}

func (s *StateDB) GetBalance(address string) *big.Int {
    s.mu.RLock()
    defer s.mu.RUnlock()
    if acc, ok := s.Accounts[address]; ok {
        return acc.Balance
    }
    return big.NewInt(0)
}

func (s *StateDB) Transfer(from, to string, amount *big.Int) bool {
    s.mu.Lock()
    defer s.mu.Unlock()

    fromAcc, ok := s.Accounts[from]
    if !ok || fromAcc.Balance.Cmp(amount) < 0 {
        return false
    }

    toAcc, ok := s.Accounts[to]
    if !ok {   If !ok {
        toAcc = &AccountState{Balance: big.NewInt(0)}
        s.Accounts[to] = toAcc
    }

    fromAcc.Balance.Sub(fromAcc.Balance, amount)
    toAcc.Balance.Add(toAcc.Balance, amount)

    return true
}

func (s *StateDB) CreateAccount(address string, balance *big.Int) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.Accounts[address] = &AccountState{
        Balance: balance,
        Nonce:   0,
    }
}


