import React, { useState, useEffect } from "react";
import axios from "axios";

function DEXInfo() {
    const [bestDEX, setBestDEX] = useState("");

    useEffect(() => {
        axios.get("http://localhost:8080/api/dex/best")
            .then(response => setBestDEX(response.data))
            .catch(error => console.error("Error fetching DEX info:", error));
    }, []);

    return (
        <div>
            <h2>Best DEX: {bestDEX}</h2>
        </div>
    );
}

export default DEXInfo;