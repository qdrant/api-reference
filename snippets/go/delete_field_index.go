package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func deleteFieldIndex() {
	client, _ := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})

	client.DeleteFieldIndex(context.Background(), &qdrant.DeleteFieldIndexCollection{
		CollectionName: "{collection_name}",
		FieldName:      "{field_name}",
	})
}
