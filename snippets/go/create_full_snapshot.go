package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func createFullSnapshot() {
	client, _ := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})

	client.CreateFullSnapshot(context.Background())
}
