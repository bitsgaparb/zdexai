import React, { useState, useEffect } from "react";
import axios from "axios";
import DEXInfo from "./components/DEXInfo";
import BridgeTokens from "./components/BridgeTokens";
import TokenPairInput from "./components/TokenPairInput";
import WalletAddressInput from "./components/WalletAddressInput";
import TransactionStatus from "./components/TransactionStatus";
import "./App.css";

function App() {
    const [tokenPair, setTokenPair] = useState("");
    const [walletAddress, setWalletAddress] = useState("");
    const [transactionStatus, setTransactionStatus] = useState("");

    useEffect(() => {
        const fetchStatus = () => {
            axios.get("http://localhost:8080/api/transaction/status")
                .then(response => setTransactionStatus(response.data))
                .catch(error => console.error("Error fetching transaction status:", error));
        };

        const interval = setInterval(fetchStatus, 5000);
        return () => clearInterval(interval);
    }, []);

    return (
        <div className="App">
            <h1>Token and DEX DApp</h1>
            <TokenPairInput onTokenPairChange={setTokenPair} />
            <WalletAddressInput onWalletAddressChange={setWalletAddress} />
            <DEXInfo />
            <BridgeTokens />
            <TransactionStatus status={transactionStatus} />
        </div>
    );
}

export default App;