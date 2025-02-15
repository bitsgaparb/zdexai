import React, { useState } from "react";

function TokenPairInput({ onTokenPairChange }) {
    const [tokenPair, setTokenPair] = useState("");

    const handleChange = (e) => {
        setTokenPair(e.target.value);
        onTokenPairChange(e.target.value);
    };

    return (
        <div>
            <input type="text" placeholder="Token Pair" value={tokenPair} onChange={handleChange} />
        </div>
    );
}

export default TokenPairInput;