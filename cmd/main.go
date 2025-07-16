package   包 main

import (   导入(
    "log"
    "net/http"
    "web3sim/api"
)

func main() {
    router := api.NewRouter()
    log.Println("Starting Web3 Sim Server on :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}


