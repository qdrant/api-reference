package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func deleteVectors() {
	client, _ := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})

	client.DeleteVectors(context.Background(), &qdrant.DeletePointVectors{
		CollectionName: "{collection_name}",
		PointsSelector: qdrant.NewPointsSelector(qdrant.NewIDNum(0), qdrant.NewIDNum(3), qdrant.NewIDNum(100)),
		Vectors: &qdrant.VectorsSelector{
			Names: []string{"text", "image"},
		},
	})
}
