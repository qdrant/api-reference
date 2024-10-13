using Qdrant.Client;
using Qdrant.Client.Grpc;

var client = new QdrantClient("localhost", 6334);

await client.SearchMatrixOffsetsAsync(
    "{collection_name}",
    filter: MatchKeyword("color", "red"),
    sample: 100,
    limit: 5
);