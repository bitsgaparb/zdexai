package api

import (
    "net/http"
    "token-dex-dapp/blockchain"
    "github.com/gorilla/mux"
)

func CreateTokenHandler(w http.ResponseWriter, r *http.Request) {
    tokenAddress, err := blockchain.CreateToken()
    if err != nil {
        http.Error(w, "Failed to create token", http.StatusInternalServerError)
        return
    }
    w.Write([]byte(tokenAddress))
}

func SetupRoutes(r *mux.Router) {
    r.HandleFunc("/api/token/create", CreateTokenHandler).Methods("POST")
}