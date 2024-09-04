package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func listSnapshots() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		panic(err)
	}

	client.ListSnapshots(context.Background(), "{collection_name}")
}