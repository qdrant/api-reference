package client

import (
	"context"
	"fmt"

	"github.com/qdrant/go-client/qdrant"
)

func facet() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		panic(err)
	}

	limit := uint64(10)
	results, err := client.Facet(context.Background(), &qdrant.FacetCounts{
		CollectionName: "{collection_name}",
		Key:            "my-payload-key",
		Limit:          &limit,
		Filter: &qdrant.Filter{
			Must: []*qdrant.Condition{
				qdrant.NewMatch("color", "red"),
			},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Query results: ", results)
}
