package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func createFieldIndex() {
	client, _ := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})

	client.CreateFieldIndex(context.Background(), &qdrant.CreateFieldIndexCollection{
		CollectionName: "{collection_name}",
		FieldName:      "name_of_the_field_to_index",
		FieldType:      qdrant.FieldType_FieldTypeKeyword.Enum(),
	})
}
