package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func updateVectors() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		panic(err)
	}

	_, err = client.UpdateVectors(context.Background(), &qdrant.UpdatePointVectors{
		CollectionName: "{collection_name}",
		Points: []*qdrant.PointVectors{
			{
				Id: qdrant.NewIDNum(1),
				Vectors: qdrant.NewVectorsMap(map[string]*qdrant.Vector{
					"image": qdrant.NewVector(0.1, 0.2, 0.3, 0.4),
				}),
			},
			{
				Id: qdrant.NewIDNum(2),
				Vectors: qdrant.NewVectorsMap(map[string]*qdrant.Vector{
					"text": qdrant.NewVector(0.9, 0.8, 0.7, 0.6, 0.5, 0.4, 0.3, 0.2),
				}),
			},
		},
	})
	if err != nil {
		panic(err)
	}
}
