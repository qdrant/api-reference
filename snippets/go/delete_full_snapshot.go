package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func deleteFullSnapshot() {
	client, _ := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})

	client.DeleteFullSnapshot(context.Background(), "{snapshot_name}")
}
