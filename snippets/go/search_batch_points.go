package client

import (
	"context"
	"fmt"

	"github.com/qdrant/go-client/qdrant"
)

func searchBatch() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		panic(err)
	}

	filter := qdrant.Filter{
		Must: []*qdrant.Condition{
			qdrant.NewMatch("city", "London"),
		},
	}
	limit := uint64(3)
	results, err := client.QueryBatch(context.Background(), &qdrant.QueryBatchPoints{
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
	if err != nil {
		panic(err)
	}
	fmt.Println("Query results: ", results)
}
