# Document as NFT
The purpose of this project is to illustrate the implementation of 'Document as NFT' concept.
Medium: https://medium.com/@megatron.web3/document-as-nft-7fa838a48737
The solidity contract after deployment on ethereum blockchain, can be interacted with as web service by the golang implementation provided.

## Steps
1. Upload the Degree.pdf to IPFS typically using a IPFS gateway like Pinata
2. Upload the nft-metadata.json to IPFS
3. Mint the nft document token using /mint API
4. Transfer document token using /transfertoken API
5. You can then use other API's like /owner and /tokenuri to verify details

## ERC721 implementation
DocNFT.sol takes the base ECR721 implementation provided by OpenZeppelin and adds any specific functionality required for 'Document as NFT'

## Generate go implementation from ERC721 contract
abigen --pkg nft --sol ~/Projects/ethereum/goland/contracts/DocNFT.sol  --out ~/Projects/ethereum/goland/contracts/documentnft.go

## Deploy the contract 
go run deploynft.go

## Start go server
go run main.go

