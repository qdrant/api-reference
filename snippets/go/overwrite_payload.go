package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func overwritePayload() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		panic(err)
	}

	client.OverwritePayload(context.Background(), &qdrant.SetPayloadPoints{
		CollectionName: "{collection_name}",
		Payload: qdrant.NewValueMap(map[string]any{
			"property1": "string",
			"property2": "string",
		}),
		PointsSelector: qdrant.NewPointsSelector(qdrant.NewIDNum(0), qdrant.NewIDNum(3), qdrant.NewIDNum(10)),
	})
}
