import React, { useState } from "react";
import axios from "axios";

function BridgeTokens() {
    const [sourceChain, setSourceChain] = useState("Solana");
    const [destinationChain, setDestinationChain] = useState("Ethereum");
    const [token, setToken] = useState("");
    const [txHash, setTxHash] = useState("");

    const handleBridge = async () => {
        const response = await axios.post("http://localhost:8080/api/bridge/tokens", {
            source: sourceChain,
            destination: destinationChain,
            token: token
        });
        setTxHash(response.data);
    };

    return (
        <div>
            <h2>Bridge Tokens</h2>
            <input type="text" placeholder="Source Chain" value={sourceChain} onChange={e => setSourceChain(e.target.value)} />
            <input type="text" placeholder="Destination Chain" value={destinationChain} onChange={e => setDestinationChain(e.target.value)} />
            <input type="text" placeholder="Token" value={token} onChange={e => setToken(e.target.value)} />
            <button onClick={handleBridge}>Bridge Tokens</button>
            {txHash && <p>Transaction Hash: {txHash}</p>}
        </div>
    );
}

export default BridgeTokens;