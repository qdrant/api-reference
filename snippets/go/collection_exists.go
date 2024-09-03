package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func collectionExists() {
	client, _ := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})

	client.CollectionExists(context.Background(), "{collection_name}")
}
