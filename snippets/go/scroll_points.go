package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func scroll() {
	client, _ := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})

	limit := uint32(1)
	client.Scroll(context.Background(), &qdrant.ScrollPoints{
		CollectionName: "{collection_name}",
		Filter: &qdrant.Filter{
			Must: []*qdrant.Condition{
				qdrant.NewMatch("color", "red"),
			},
		},
		Limit:       &limit,
		WithPayload: qdrant.NewWithPayloadEnable(true),
		WithVectors: qdrant.NewWithVectorsEnable(false),
	})
}
