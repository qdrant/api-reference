package client

import (
	"context"
	"fmt"

	"github.com/qdrant/go-client/qdrant"
)

func listFullSnapshots() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		panic(err)
	}

	snapshots, err := client.ListFullSnapshots(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("Full snapshots: ", snapshots)
}
