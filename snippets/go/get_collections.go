package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func listCollections() {
	client, _ := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})

	client.ListCollections(context.Background())
}
