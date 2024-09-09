package client

import (
	"context"
	"fmt"

	"github.com/qdrant/go-client/qdrant"
)

func searchGroups() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		panic(err)
	}

	groupSize := uint64(2)
	results, err := client.QueryGroups(context.Background(), &qdrant.QueryPointGroups{
		CollectionName: "{collection_name}",
		Query:          qdrant.NewQuery(0.01, 0.45, 0.67),
		GroupBy:        "document_id",
		GroupSize:      &groupSize,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Query results: ", results)
}
