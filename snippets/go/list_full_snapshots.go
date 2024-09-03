package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func listFullSnapshots() {
	client, _ := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})

	client.ListFullSnapshots(context.Background())
}
