package client

import (
	"context"
	"fmt"

	"github.com/qdrant/go-client/qdrant"
)

func count() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		panic(err)
	}

	count, err := client.Count(context.Background(), &qdrant.CountPoints{
		CollectionName: "{collection_name}",
		Filter: &qdrant.Filter{
			Must: []*qdrant.Condition{
				qdrant.NewMatch("color", "red"),
			},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Count:", count)
}
