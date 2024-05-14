using Qdrant.Client;
using Qdrant.Client.Grpc;

var client = new QdrantClient("localhost", 6334);

await client.DiscoverAsync(
  collectionName: "{collection_name}",
  target: new TargetVector
  {
    Single = new VectorExample { Vector = new float[] { 0.2f, 0.1f, 0.9f, 0.7f }, }
  },
  context:
  [
    new()
    {
      Positive = new VectorExample { Id = 100 },
      Negative = new VectorExample { Id = 718 }
    },
    new()
    {
      Positive = new VectorExample { Id = 200 },
      Negative = new VectorExample { Id = 300 }
    }
  ],
  limit: 10
);
