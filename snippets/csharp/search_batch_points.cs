using Qdrant.Client;
using Qdrant.Client.Grpc;
using static Qdrant.Client.Grpc.Conditions;

var client = new QdrantClient("localhost", 6334);

var filter = MatchKeyword("city", "London");

var searches = new List<SearchPoints>
{
  new()
  {
    Vector = { new float[] { 0.2f, 0.1f, 0.9f, 0.7f } },
    Filter = filter,
    Limit = 3
  },
  new()
  {
    Vector = { new float[] { 0.5f, 0.3f, 0.2f, 0.3f } },
    Filter = filter,
    Limit = 3
  }
};

await client.SearchBatchAsync(collectionName: "{collection_name}", searches: searches);
