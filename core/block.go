package   包 core

type Block struct {
    Number       int
    Transactions []*Transaction
    PrevHash     string
    Hash         string
    Timestamp    int64
}


