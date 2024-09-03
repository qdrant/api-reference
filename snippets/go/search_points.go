package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func search() {
	client, _ := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})

	limit := uint64(3)
	client.Query(context.Background(), &qdrant.QueryPoints{
		CollectionName: "{collection_name}",
		Query:          qdrant.NewQuery(0.2, 0.1, 0.9, 0.7),
		Filter: &qdrant.Filter{
			Must: []*qdrant.Condition{
				qdrant.NewMatch("city", "London"),
			},
		},
		Limit: &limit,
	})
}
