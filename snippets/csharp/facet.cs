using Qdrant.Client;
using Qdrant.Client.Grpc;

var client = new QdrantClient("localhost", 6334);

await client.FacetAsync(
	"{collection_name}",
	filter: MatchKeyword("color", "red"),
	key: "my-payload-key",
	limit: 10
);