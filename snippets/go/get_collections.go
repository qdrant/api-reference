package client

import (
	"context"
	"fmt"

	"github.com/qdrant/go-client/qdrant"
)

func listCollections() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		panic(err)
	}

	collections, err := client.ListCollections(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("Collections: ", collections)
}
