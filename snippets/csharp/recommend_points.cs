using Qdrant.Client;
using Qdrant.Client.Grpc;
using static Qdrant.Client.Grpc.Conditions;

var client = new QdrantClient("localhost", 6334);

await client.RecommendAsync(
    "{collection_name}",
    positive: new ulong[] { 100, 231 },
    negative: new ulong[] { 718 },
    filter: MatchKeyword("city", "London"),
    limit: 3
);
