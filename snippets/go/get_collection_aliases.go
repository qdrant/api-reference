package client

import (
	"context"
	"fmt"

	"github.com/qdrant/go-client/qdrant"
)

func listCollectionAliases() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		panic(err)
	}

	aliases, err := client.ListCollectionAliases(context.Background(), "{collection_name}")
	if err != nil {
		panic(err)
	}
	fmt.Println("Collection aliases: ", aliases)
}
