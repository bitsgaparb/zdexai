package blockchain

import (
    "context"
    "github.com/solana/go-solana-sdk/client"
    "github.com/solana/go-solana-sdk/program/token"
    "github.com/solana/go-solana-sdk/types"
    "github.com/solana/go-solana-sdk/rpc"
)

func CreateToken() (string, error) {
    cli := rpc.NewClient("https://api.mainnet-beta.solana.com")
    payerAccount := types.AccountFromPrivateKeyBytes([]byte("your-private-key"))
    mintAccount := types.NewAccount()

    tokenProgram := token.NewProgram()
    instruction := tokenProgram.InitializeMint(
        mintAccount.PublicKey(),
        6, // Decimals
        payerAccount.PublicKey(),
        nil, // Freeze Authority
    )

    _, err := cli.SendTransaction([]types.Instruction{instruction}, []types.Account{payerAccount, mintAccount})
    if err != nil {
        return "", err
    }

    return mintAccount.PublicKey().String(), nil
}