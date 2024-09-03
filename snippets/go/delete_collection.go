package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func deleteCollection() {
	client, _ := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})

	client.DeleteCollection(context.Background(), "{collection_name}")
}
