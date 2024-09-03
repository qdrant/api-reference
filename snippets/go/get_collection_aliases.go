package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func listCollectionAliases() {
	client, _ := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})

	client.ListCollectionAliases(context.Background(), "{collection_name}")
}
