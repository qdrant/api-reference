package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func listAliases() {
	client, _ := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})

	client.ListAliases(context.Background())
}
