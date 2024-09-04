package client

import (
	"context"

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

	client.ListAliases(context.Background())
}