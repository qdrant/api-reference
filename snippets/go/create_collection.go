package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func createCollection() {
	client, _ := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})

	client.CreateCollection(context.Background(), &qdrant.CreateCollection{
		CollectionName: "{collection_name}",
		VectorsConfig: qdrant.NewVectorsConfig(&qdrant.VectorParams{
			Size:     100,
			Distance: qdrant.Distance_Cosine,
		}),
	})
}
