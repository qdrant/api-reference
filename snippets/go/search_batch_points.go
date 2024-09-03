package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func searchBatch() {
	client, _ := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})

	filter := qdrant.Filter{
		Must: []*qdrant.Condition{
			qdrant.NewMatch("city", "London"),
		},
	}
	limit := uint64(3)
	client.QueryBatch(context.Background(), &qdrant.QueryBatchPoints{
		CollectionName: "{collection_name}",
		QueryPoints: []*qdrant.QueryPoints{
			{
				CollectionName: "{collection_name}",
				Query:          qdrant.NewQuery(0.2, 0.1, 0.9, 0.7),
				Limit:          &limit,
				Filter:         &filter,
			},
			{
				CollectionName: "{collection_name}",
				Query:          qdrant.NewQuery(0.5, 0.3, 0.2, 0.3),
				Limit:          &limit,
				Filter:         &filter,
			},
		},
	})
}
