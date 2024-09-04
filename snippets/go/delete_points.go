package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func delete() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		panic(err)
	}

	_, err = client.Delete(context.Background(), &qdrant.DeletePoints{
		CollectionName: "{collection_name}",
		Points: qdrant.NewPointsSelectorIDs([]*qdrant.PointId{
			qdrant.NewIDNum(0), qdrant.NewIDNum(3), qdrant.NewIDNum(100),
		}),
	})
	if err != nil {
		panic(err)
	}

	_, err = client.Delete(context.Background(), &qdrant.DeletePoints{
		CollectionName: "{collection_name}",
		Points: qdrant.NewPointsSelectorFilter(&qdrant.Filter{
			Must: []*qdrant.Condition{
				qdrant.NewMatch("color", "red"),
			},
		}),
	})
	if err != nil {
		panic(err)
	}
}
