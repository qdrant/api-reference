using Qdrant.Client;
using Qdrant.Client.Grpc;

var client = new QdrantClient("localhost", 6334);

await client.UpdateBatchAsync(
    "{collection_name}",
    [
        new()
        {
            Upsert = new()
            {
                Points =
                {
                    new PointStruct { Id = 1, Vectors = new[] { 0.9f, 0.1f, 0.1f } },
                }
            }
        },
        new()
        {
            UpdateVectors = new()
            {
                Points =
                {
                    new PointVectors { Id = 1, Vectors = new[] { 0.9f, 0.1f, 0.1f } },
                }
            }
        },
        new()
        {
            SetPayload = new()
            {
                PointsSelector = new PointsSelector { Points = new PointsIdsList { Ids = { 1 } } },
                Payload = { ["test_payload_2"] = 2, ["test_payload_3"] = 3 }
            }
        }
    ]
);
