import React, { useState } from "react";

function WalletAddressInput({ onWalletAddressChange }) {
    const [walletAddress, setWalletAddress] = useState("");

    const handleChange = (e) => {
        setWalletAddress(e.target.value);
        onWalletAddressChange(e.target.value);
    };

    return (
        <div>
            <input type="text" placeholder="Wallet Address" value={walletAddress} onChange={handleChange} />
        </div>
    );
}

export default WalletAddressInput;