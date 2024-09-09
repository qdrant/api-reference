package client

import (
	"context"
	"fmt"

	"github.com/qdrant/go-client/qdrant"
)

func createSnapshot() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		panic(err)
	}

	snapshot, err := client.CreateSnapshot(context.Background(), "{collection_name}")
	if err != nil {
		panic(err)
	}
	fmt.Println("Snapshot created: ", snapshot.Name)
}
