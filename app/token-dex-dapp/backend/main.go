package main

import (
    "log"
    "net/http"
    "token-dex-dapp/api"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    api.SetupRoutes(r)
    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}