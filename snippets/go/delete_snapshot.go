package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func deleteSnapshot() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		panic(err)
	}

	err = client.DeleteSnapshot(context.Background(), "{collection_name}", "{snapshot_name}")
	if err != nil {
		panic(err)
	}
}
