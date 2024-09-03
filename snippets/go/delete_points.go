package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func delete() {
	client, _ := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})

	client.Delete(context.Background(), &qdrant.DeletePoints{
		CollectionName: "{collection_name}",
		Points: qdrant.NewPointsSelectorIDs([]*qdrant.PointId{
			qdrant.NewIDNum(0), qdrant.NewIDNum(3), qdrant.NewIDNum(100),
		}),
	})

	client.Delete(context.Background(), &qdrant.DeletePoints{
		CollectionName: "{collection_name}",
		Points: qdrant.NewPointsSelectorFilter(&qdrant.Filter{
			Must: []*qdrant.Condition{
				qdrant.NewMatch("color", "red"),
			},
		}),
	})
}
