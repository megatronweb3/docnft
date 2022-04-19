package handler

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	Models "goland/models"
	Modules "goland/modules"
	"net/http"
)

// ClientHandler ethereum client instance
type ClientHandler struct {
	*ethclient.Client
}

func (client ClientHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Get parameter from url request
	vars := mux.Vars(r)
	action := vars["action"]

	// Get the query parameters from url request
	//address := r.URL.Query().Get("address")
	//hash := r.URL.Query().Get("hash")

	// Set our response header
	w.Header().Set("Content-Type", "application/json")

	// Handle each request using the module parameter:
	switch action {
	case "symbol":
		_address := Modules.GetTokenSymbol(*client.Client)
		json.NewEncoder(w).Encode(_address)
		return

	case "name":
		_address := Modules.GetTokenName(*client.Client)
		json.NewEncoder(w).Encode(_address)
		return

	case "owner":
		decoder := json.NewDecoder(r.Body)
		var t Models.TokenId

		err := decoder.Decode(&t)
		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(&Models.Error{
				Code:    400,
				Message: "Malformed request",
			})
			return
		}
		_ownerAddress := Modules.GetOwner(*client.Client, t.TokenId)
		json.NewEncoder(w).Encode(_ownerAddress)

		return

	case "tokenuri":
		decoder := json.NewDecoder(r.Body)
		var t Models.TokenId

		err := decoder.Decode(&t)
		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(&Models.Error{
				Code:    400,
				Message: "Malformed request",
			})
			return
		}
		_ownerAddress := Modules.GetTokenURI(*client.Client, t.TokenId)
		json.NewEncoder(w).Encode(_ownerAddress)

		return

	case "tokentransfer":
		decoder := json.NewDecoder(r.Body)
		var t Models.TokenTransferRequest

		err := decoder.Decode(&t)
		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(&Models.Error{
				Code:    400,
				Message: "Malformed request",
			})
			return
		}
		_transaction := Modules.TransferToken(*client.Client, t.PrivateKey, t.FromAddress, t.ToAddress, t.TokenId)
		json.NewEncoder(w).Encode(_transaction)

		return

	case "mint":
		decoder := json.NewDecoder(r.Body)
		var t Models.MintTokenRequest

		err := decoder.Decode(&t)
		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(&Models.Error{
				Code:    400,
				Message: "Malformed request",
			})
			return
		}
		_transaction := Modules.Mint(*client.Client, t.PrivateKey, t.PublicAddress, t.TokenURI)
		json.NewEncoder(w).Encode(_transaction)

		return

	}

}
