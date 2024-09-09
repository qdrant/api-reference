package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func deleteFieldIndex() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		panic(err)
	}

	_, err = client.DeleteFieldIndex(context.Background(), &qdrant.DeleteFieldIndexCollection{
		CollectionName: "{collection_name}",
		FieldName:      "{field_name}",
	})
	if err != nil {
		panic(err)
	}
}
