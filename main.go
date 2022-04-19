package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	Handlers "goland/handler"
	Modules "goland/modules"
	"log"
	"net/http"
)

func main() {
	// Create a client instance to connect to our provider
	client, err := ethclient.Dial(Modules.Config().Network)

	if err != nil {
		fmt.Println(err)
	}

	// Create a mux router
	r := mux.NewRouter()

	// We will define a single endpoint
	r.Handle("/api/v1/nft/{action}", Handlers.ClientHandler{client})
	log.Fatal(http.ListenAndServe(":8080", r))
}
