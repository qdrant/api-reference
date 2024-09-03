package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func overwritePayload() {
	client, _ := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})

	client.OverwritePayload(context.Background(), &qdrant.SetPayloadPoints{
		CollectionName: "{collection_name}",
		Payload: qdrant.NewValueMap(map[string]any{
			"property1": "string",
			"property2": "string",
		}),
		PointsSelector: qdrant.NewPointsSelector(qdrant.NewIDNum(0), qdrant.NewIDNum(3), qdrant.NewIDNum(10)),
	})
}
