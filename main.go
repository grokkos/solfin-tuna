package main

import (
	"context"
	"fmt"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/rpc"
)

func main() {
	c := client.NewClient(rpc.MainnetRPCEndpoint)

	response, err := c.GetVersion(context.TODO())

	if err != nil {
		panic(err)
	}

	fmt.Println("Version", response.SolanaCore)
}
