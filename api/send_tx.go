package   包 api

import (   导入(
    "encoding/json"
    "math/big"
    "net/http"
    "web3sim/core"
    "web3sim/db"
)

var state = db.NewStateDB()
var engine = core.ContractEngine{State: state}

func init() {
    state.CreateAccount("0xabc", big.NewInt(1000000000000000000))
    state.CreateAccount("0xdef", big.NewInt(0))
}

type TxInput struct {
    From   string `json:"from"`
    To     string `json:"to"`
    Amount string `json:"amount"`
}

func SendTxHandler(w http.ResponseWriter, r *http.Request) {
    var input TxInput
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    amt, ok := new(big.Int).SetString(input.Amount, 10)
    if !ok {
        http.Error(w, "Invalid amount", http.StatusBadRequest)
        return
    }

    tx := core.NewTransaction(input.From, input.To, amt)
    success := engine.ExecuteTransfer(tx)

    if success {
        json.NewEncoder(w).Encode(map[string]string{
            "message": "✅ 交易成功！",
        })
    } else {
        json.NewEncoder(w).Encode(map[string]string{
            "error": "❌ 交易失败，余额不足或账户不存在。",
        })
    }
}


