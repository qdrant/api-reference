using Qdrant.Client;
using static Qdrant.Client.Grpc.Conditions;

var client = new QdrantClient("localhost", 6334);

await client.RecommendGroupsAsync(
    "{collection_name}",
    "document_id",
    groupSize: 3,
    positive: new ulong[] { 100, 231 },
    negative: new ulong[] { 718 },
    filter: MatchKeyword("city", "London"),
    limit: 3
);
