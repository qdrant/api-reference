package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func updateCollection() {
	client, _ := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})

	threshold := uint64(10000)
	client.UpdateCollection(context.Background(), &qdrant.UpdateCollection{
		CollectionName: "{collection_name}",
		OptimizersConfig: &qdrant.OptimizersConfigDiff{
			IndexingThreshold: &threshold,
		},
	})
}
