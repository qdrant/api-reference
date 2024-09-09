package client

import (
	"context"
	"fmt"

	"github.com/qdrant/go-client/qdrant"
)

func getCollection() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		panic(err)
	}

	info, err := client.GetCollectionInfo(context.Background(), "{collection_name}")
	if err != nil {
		panic(err)
	}
	fmt.Println("Collection info: ", info)
}
