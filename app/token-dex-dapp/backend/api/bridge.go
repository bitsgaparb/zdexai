package api

import (
    "net/http"
    "token-dex-dapp/bridge"
    "github.com/gorilla/mux"
)

func BridgeTokensHandler(w http.ResponseWriter, r *http.Request) {
    sourceChain := r.URL.Query().Get("source")
    destinationChain := r.URL.Query().Get("destination")
    token := r.URL.Query().Get("token")

    txHash, err := bridge.BridgeTokens(sourceChain, destinationChain, token)
    if err != nil {
        http.Error(w, "Failed to bridge tokens", http.StatusInternalServerError)
        return
    }
    w.Write([]byte(txHash))
}

func SetupRoutes(r *mux.Router) {
    r.HandleFunc("/api/bridge/tokens", BridgeTokensHandler).Methods("POST")
}