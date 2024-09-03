package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func listSnapshots() {
	client, _ := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})

	client.ListSnapshots(context.Background(), "{collection_name}")
}
