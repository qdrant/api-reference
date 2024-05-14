using Qdrant.Client;
using static Qdrant.Client.Grpc.Conditions;

var client = new QdrantClient("localhost", 6334);

await client.ScrollAsync(
  collectionName: "{collection_name}",
  filter: MatchKeyword("color", "red"),
  limit: 1,
  payloadSelector: true
);
