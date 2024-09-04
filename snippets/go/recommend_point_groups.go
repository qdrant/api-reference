package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func recommendGroups() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		panic(err)
	}

	groupSize := uint64(5)
	client.QueryGroups(context.Background(), &qdrant.QueryPointGroups{
		CollectionName: "{collection_name}",
		Query: qdrant.NewQueryRecommend(&qdrant.RecommendInput{
			Positive: []*qdrant.VectorInput{
				qdrant.NewVectorInputID(qdrant.NewIDNum(100)),
				qdrant.NewVectorInputID(qdrant.NewIDNum(231)),
			},
			Negative: []*qdrant.VectorInput{
				qdrant.NewVectorInputID(qdrant.NewIDNum(718)),
			},
		}),
		GroupBy:   "document_id",
		GroupSize: &groupSize,
	})
}
