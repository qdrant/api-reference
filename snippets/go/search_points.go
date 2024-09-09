package client

import (
	"context"
	"fmt"

	"github.com/qdrant/go-client/qdrant"
)

func search() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		panic(err)
	}

	limit := uint64(3)
	results, err := client.Query(context.Background(), &qdrant.QueryPoints{
		CollectionName: "{collection_name}",
		Query:          qdrant.NewQuery(0.2, 0.1, 0.9, 0.7),
		Filter: &qdrant.Filter{
			Must: []*qdrant.Condition{
				qdrant.NewMatch("city", "London"),
			},
		},
		Limit: &limit,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Results: ", results)
}
