package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	Modules "goland/modules"
)

func main() {

	// Create a client instance to connect to our provider
	client, err := ethclient.Dial(Modules.Config().Network)

	if err != nil {
		fmt.Println(err)
	}

	privateKey := "b6968a12942e21ea2f56bc38171c6abc026367bb9a3739c648754c4f9c83bce3"
	address := Modules.DeployContract(client, privateKey)

	fmt.Println("ContractAddress=", *address)
}
