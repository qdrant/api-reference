using Qdrant.Client;
using Qdrant.Client.Grpc;
using static Qdrant.Client.Grpc.Conditions;

var client = new QdrantClient("localhost", 6334);

var filter = MatchKeyword("city", "london");

await client.RecommendBatchAsync(
  collectionName: "{collection_name}",
  recommendSearches:
  [
    new()
    {
      CollectionName = "{collection_name}",
      Positive = { new PointId[] { 100, 231 } },
      Negative = { new PointId[] { 718 } },
      Limit = 3,
      Filter = filter,
    },
    new()
    {
      CollectionName = "{collection_name}",
      Positive = { new PointId[] { 200, 67 } },
      Negative = { new PointId[] { 300 } },
      Limit = 3,
      Filter = filter,
    }
  ]
);
