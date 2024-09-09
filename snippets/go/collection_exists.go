package client

import (
	"context"
	"fmt"

	"github.com/qdrant/go-client/qdrant"
)

func collectionExists() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		panic(err)
	}

	exists, err := client.CollectionExists(context.Background(), "{collection_name}")
	if err != nil {
		panic(err)
	}
	fmt.Println("Collection exists: ", exists)
}
