package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func batchUpdate() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		panic(err)
	}

	_, err = client.UpdateBatch(context.Background(), &qdrant.UpdateBatchPoints{
		CollectionName: "{collection_name}",
		Operations: []*qdrant.PointsUpdateOperation{
			qdrant.NewPointsUpdateUpsert(&qdrant.PointsUpdateOperation_PointStructList{
				Points: []*qdrant.PointStruct{{
					Id:      qdrant.NewIDNum(1),
					Vectors: qdrant.NewVectors(1.0, 2.0, 3.0, 4.0),
				}},
			}),
			qdrant.NewPointsUpdateUpdateVectors(&qdrant.PointsUpdateOperation_UpdateVectors{
				Points: []*qdrant.PointVectors{
					{
						Id:      qdrant.NewIDNum(1),
						Vectors: qdrant.NewVectors(0.1, 0.2, 0.3, 0.4),
					},
				},
			}),
			qdrant.NewPointsUpdateSetPayload(&qdrant.PointsUpdateOperation_SetPayload{
				PointsSelector: qdrant.NewPointsSelector(qdrant.NewIDNum(1)),
				Payload: qdrant.NewValueMap(map[string]any{
					"test_payload_2": 2,
					"test_payload_3": 3,
				}),
			}),
		},
	})
	if err != nil {
		panic(err)
	}
}
