package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func deleteCollection() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		panic(err)
	}

	err = client.DeleteCollection(context.Background(), "{collection_name}")
	if err != nil {
		panic(err)
	}
}
