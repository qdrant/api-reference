using Qdrant.Client;
using Qdrant.Client.Grpc;

var client = new QdrantClient("localhost", 6334);

await client.QueryBatchAsync(
	collectionName: "{collection_name}",
	queries: new List<QueryPoints>
	{
		new() { Query = new float[] { 0.1f, 0.2f, 0.3f }, },
		new() { Query = new float[] { 0.4f, 0.5f, 0.6f }, }
	}
);
