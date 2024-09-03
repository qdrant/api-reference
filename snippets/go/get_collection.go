package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func getCollection() {
	client, _ := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})

	client.GetCollection(context.Background(), "{collection_name}")
}
