package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func updateAlias() {
	client, _ := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})

	client.CreateAlias(context.Background(), "production_collection", "example_collection")

	client.DeleteAlias(context.Background(), "production_collection")

	client.RenameAlias(context.Background(), "production_collection", "legacy_collection")
}
