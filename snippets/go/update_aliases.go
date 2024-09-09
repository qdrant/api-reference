package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func updateAlias() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		panic(err)
	}

	err = client.CreateAlias(context.Background(), "production_collection", "example_collection")
	if err != nil {
		panic(err)
	}

	err = client.DeleteAlias(context.Background(), "production_collection")
	if err != nil {
		panic(err)
	}

	err = client.RenameAlias(context.Background(), "production_collection", "legacy_collection")
	if err != nil {
		panic(err)
	}
}
