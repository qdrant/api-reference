package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func createSnapshot() {
	client, _ := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})

	client.CreateSnapshot(context.Background(), "{collection_name}")
}
