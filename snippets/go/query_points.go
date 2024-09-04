package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func query() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		panic(err)
	}

	// Query nearest by ID
	client.Query(context.Background(), &qdrant.QueryPoints{
		CollectionName: "{collection_name}",
		Query:          qdrant.NewQueryID(qdrant.NewID("43cf51e2-8777-4f52-bc74-c2cbde0c8b04")),
	})

	// Recommend on the average of these vectors
	client.Query(context.Background(), &qdrant.QueryPoints{
		CollectionName: "{collection_name}",
		Query: qdrant.NewQueryRecommend(&qdrant.RecommendInput{
			Positive: []*qdrant.VectorInput{
				qdrant.NewVectorInputID(qdrant.NewID("43cf51e2-8777-4f52-bc74-c2cbde0c8b04")),
				qdrant.NewVectorInput(0.11, 0.35, 0.6),
			},
			Negative: []*qdrant.VectorInput{
				qdrant.NewVectorInput(0.01, 0.45, 0.67),
			},
		}),
	})

	// Fusion query
	client.Query(context.Background(), &qdrant.QueryPoints{
		CollectionName: "{collection_name}",
		Prefetch: []*qdrant.PrefetchQuery{
			{
				Query: qdrant.NewQuerySparse([]uint32{1, 42}, []float32{0.22, 0.8}),
				Using: qdrant.PtrOf("sparse"),
			},
			{
				Query: qdrant.NewQuery(0.01, 0.45, 0.67),
				Using: qdrant.PtrOf("dense"),
			},
		},
		Query: qdrant.NewQueryFusion(qdrant.Fusion_RRF),
	})

	// 2-stage query
	client.Query(context.Background(), &qdrant.QueryPoints{
		CollectionName: "{collection_name}",
		Prefetch: []*qdrant.PrefetchQuery{
			{
				Query: qdrant.NewQuery(0.01, 0.45, 0.67),
			},
		},
		Query: qdrant.NewQueryMulti([][]float32{
			{0.1, 0.2},
			{0.2, 0.1},
			{0.8, 0.9},
		}),
		Using: qdrant.PtrOf("colbert"),
	})

	// Random sampling (as of 1.11.0)
	client.Query(context.Background(), &qdrant.QueryPoints{
		CollectionName: "{collection_name}",
		Query:          qdrant.NewQuerySample(qdrant.Sample_Random),
	})
}
