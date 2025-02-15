package ai

import (
    "log"
    "time"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

type DEXRate struct {
    Exchange string  `json:"exchange"`
    Rate     float64 `json:"rate"`
}

func MonitorDEXRates() (string, error) {
    resp, err := http.Get("https://api.dexplatform.com/rates")
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    var rates []DEXRate
    if err := json.Unmarshal(body, &rates); err != nil {
        return "", err
    }

    var bestRate DEXRate
    for _, rate := range rates {
        if rate.Rate > bestRate.Rate {
            bestRate = rate
        }
    }

    return bestRate.Exchange, nil
}