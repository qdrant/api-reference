package client

import (
	"context"
	"fmt"

	"github.com/qdrant/go-client/qdrant"
)

func listAliases() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		panic(err)
	}

	aliases, err := client.ListAliases(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("Aliases: ", aliases)
}
