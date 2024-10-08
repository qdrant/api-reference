package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func clearPayload() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		panic(err)
	}

	_, err = client.ClearPayload(context.Background(), &qdrant.ClearPayloadPoints{
		CollectionName: "{collection_name}",
		Points:         qdrant.NewPointsSelector(qdrant.NewIDNum(0), qdrant.NewID("3"), qdrant.NewIDNum(100)),
	})
	if err != nil {
		panic(err)
	}
}
