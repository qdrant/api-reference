package client

import (
	"context"
	"fmt"

	"github.com/qdrant/go-client/qdrant"
)

func searchMatiOffsets() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		panic(err)
	}

	sample := uint64(100)
	limit := uint64(5)
	results, err := client.SearchMatrixOffsets(context.Background(), &qdrant.SearchMatrixPoints{
		CollectionName: "{collection_name}",
		Sample:         &sample,
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
