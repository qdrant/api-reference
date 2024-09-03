package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func recommendBatch() {
	client, _ := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})

	client.QueryBatch(context.Background(), &qdrant.QueryBatchPoints{
		CollectionName: "{collection_name}",
		QueryPoints: []*qdrant.QueryPoints{
			{
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
			},
			{
				CollectionName: "{collection_name}",
				Query: qdrant.NewQueryRecommend(&qdrant.RecommendInput{
					Positive: []*qdrant.VectorInput{
						qdrant.NewVectorInputID(qdrant.NewIDNum(200)),
						qdrant.NewVectorInputID(qdrant.NewIDNum(67)),
					},
					Negative: []*qdrant.VectorInput{
						qdrant.NewVectorInputID(qdrant.NewIDNum(300)),
					},
				}),
			},
		},
	})
}
