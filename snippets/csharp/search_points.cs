using Qdrant.Client;
using Qdrant.Client.Grpc;
using static Qdrant.Client.Grpc.Conditions;

var client = new QdrantClient("localhost", 6334);

await client.SearchAsync(
	collectionName: "{collection_name}",
	vector: new float[] { 0.2f, 0.1f, 0.9f, 0.7f },
	filter: MatchKeyword("city", "London"),
	searchParams: new SearchParams { Exact = false, HnswEf = 128 },
	limit: 3
);