package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func deletePayload() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		panic(err)
	}

	_, err = client.DeletePayload(context.Background(), &qdrant.DeletePayloadPoints{
		CollectionName: "{collection_name}",
		PointsSelector: qdrant.NewPointsSelector(qdrant.NewIDNum(0), qdrant.NewIDNum(3), qdrant.NewIDNum(100)),
		Keys:           []string{"color", "price"},
	})
	if err != nil {
		panic(err)
	}
}
