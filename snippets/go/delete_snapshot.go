package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func deleteSnapshot() {
	client, _ := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})

	client.DeleteSnapshot(context.Background(), "{collection_name}", "{snapshot_name}")
}
