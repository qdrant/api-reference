package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func updateCollection() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		panic(err)
	}

	threshold := uint64(10000)
	client.UpdateCollection(context.Background(), &qdrant.UpdateCollection{
		CollectionName: "{collection_name}",
		OptimizersConfig: &qdrant.OptimizersConfigDiff{
			IndexingThreshold: &threshold,
		},
	})
}
