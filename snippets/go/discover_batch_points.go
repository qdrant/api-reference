package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func discoverBatch() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		panic(err)
	}

	client.QueryBatch(context.Background(), &qdrant.QueryBatchPoints{
		CollectionName: "{collection_name}",
		QueryPoints: []*qdrant.QueryPoints{
			{
				CollectionName: "{collection_name}",
				Query: qdrant.NewQueryDiscover(&qdrant.DiscoverInput{
					Target: qdrant.NewVectorInput(0.2, 0.1, 0.9, 0.7),
					Context: &qdrant.ContextInput{
						Pairs: []*qdrant.ContextInputPair{{
							Positive: qdrant.NewVectorInputID(qdrant.NewIDNum(100)),
							Negative: qdrant.NewVectorInputID(qdrant.NewIDNum(718)),
						}, {
							Positive: qdrant.NewVectorInputID(qdrant.NewIDNum(200)),
							Negative: qdrant.NewVectorInputID(qdrant.NewIDNum(300)),
						}},
					},
				}),
			},
			{
				CollectionName: "{collection_name}",
				Query: qdrant.NewQueryDiscover(&qdrant.DiscoverInput{
					Target: qdrant.NewVectorInput(0.5, 0.3, 0.2, 0.3),
					Context: &qdrant.ContextInput{
						Pairs: []*qdrant.ContextInputPair{{
							Positive: qdrant.NewVectorInputID(qdrant.NewIDNum(342)),
							Negative: qdrant.NewVectorInputID(qdrant.NewIDNum(213)),
						}, {
							Positive: qdrant.NewVectorInputID(qdrant.NewIDNum(100)),
							Negative: qdrant.NewVectorInputID(qdrant.NewIDNum(200)),
						}},
					},
				}),
			},
		},
	})
}
