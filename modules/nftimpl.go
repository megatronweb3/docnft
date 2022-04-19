package modules

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"goland/contracts"
	Models "goland/models"
	"math/big"
)

// Get address of the owner of the token
func GetOwner(client ethclient.Client, tokenId int64) *Models.Address {
	txn, err := getConn(&client).OwnerOf(&bind.CallOpts{}, big.NewInt(tokenId))
	if err != nil {
		panic(err)
	}
	return &Models.Address{
		Address: txn.String(),
	}
}

// Get name of the token
func GetTokenName(client ethclient.Client) *Models.TokenName {
	txn, err := getConn(&client).Name(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	return &Models.TokenName{
		TokenName: txn,
	}
}

// Get symbol of the token
func GetTokenSymbol(client ethclient.Client) *Models.TokenSymbol {
	txn, err := getConn(&client).Symbol(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	return &Models.TokenSymbol{
		TokenSymbol: txn,
	}
}

// Get IPFS token URL
func GetTokenURI(client ethclient.Client, tokenId int64) *Models.TokenURI {
	txn, err := getConn(&client).TokenURI(&bind.CallOpts{}, big.NewInt(tokenId))
	if err != nil {
		panic(err)
	}
	return &Models.TokenURI{
		TokenURI: txn,
	}
}

//Transfer token
func TransferToken(client ethclient.Client, privateKeyAddress string, fromAddress string, toAddress string, tokenId int64) *Models.Transaction {
	tx, err := getConn(&client).SafeTransferFrom0(getAccountAuth(&client, privateKeyAddress), common.HexToAddress(fromAddress),
		common.HexToAddress(toAddress), big.NewInt(tokenId), nil)

	if err != nil {
		panic(err)
	}
	return &Models.Transaction{
		Hash:     tx.Hash().String(),
		Value:    tx.Value().String(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice().Uint64(),
		To:       tx.To().String(),
		Nonce:    tx.Nonce(),
	}
}

//Mint NFT
func Mint(client ethclient.Client, privateKeyAddress string, publicAddress string, tokenURI string) *Models.Transaction {
	tx, err := getConn(&client).MintNFT(getAccountAuth(&client, privateKeyAddress), common.HexToAddress(publicAddress), tokenURI)
	if err != nil {
		panic(err)
	}
	return &Models.Transaction{
		Hash:     tx.Hash().String(),
		Value:    tx.Value().String(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice().Uint64(),
		To:       tx.To().String(),
		Nonce:    tx.Nonce(),
	}
}

func DeployContract(client *ethclient.Client, privateKey string) *Models.Address {
	address, _, _, err := nft.DeployDocNFT(getAccountAuth(client, privateKey), client)

	if err != nil {
		panic(err)
	}

	return &Models.Address{
		Address: address.String(),
	}
}

func getConn(client *ethclient.Client) *nft.DocNFT {
	//Address where contract is deployed
	contractAddress := Config().ContractAddress

	// create auth and transaction package for deploying smart contract
	conn, err := nft.NewDocNFT(common.HexToAddress(contractAddress), client)
	if err != nil {
		panic(err)
	}

	return conn
}

func getAccountAuth(client *ethclient.Client, privateKeyAddress string) *bind.TransactOpts {

	privateKey, err := crypto.HexToECDSA(privateKeyAddress)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}
	fmt.Println("nounce=", nonce)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = big.NewInt(1000000)

	return auth
}
