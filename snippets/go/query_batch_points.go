package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func queryBatch() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		panic(err)
	}

	client.QueryBatch(context.Background(), &qdrant.QueryBatchPoints{
		CollectionName: "{collection_name}",
		QueryPoints: []*qdrant.QueryPoints{
			{
				CollectionName: "{collection_name}",
				Query:          qdrant.NewQuery(0.01, 0.45, 0.67),
			},
			{
				CollectionName: "{collection_name}",
				Query:          qdrant.NewQuery(0.11, 0.35, 0.6),
			},
		},
	})
}
