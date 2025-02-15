package api

import (
    "net/http"
    "token-dex-dapp/ai"
    "github.com/gorilla/mux"
)

func BestDEXHandler(w http.ResponseWriter, r *http.Request) {
    bestDEX, err := ai.MonitorDEXRates()
    if err != nil {
        http.Error(w, "Failed to monitor DEX rates", http.StatusInternalServerError)
        return
    }
    w.Write([]byte(bestDEX))
}

func SetupRoutes(r *mux.Router) {
    r.HandleFunc("/api/dex/best", BestDEXHandler).Methods("GET")
}