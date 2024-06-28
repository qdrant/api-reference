using Qdrant.Client;
using Qdrant.Client.Grpc;

var client = new QdrantClient("localhost", 6334);

// Query nearest by ID
await client.QueryAsync(
	collectionName: "{collection_name}",
	query: Guid.Parse("43cf51e2-8777-4f52-bc74-c2cbde0c8b04")
);

// Recommend on the average of these vectors
await client.QueryAsync(
	collectionName: "{collection_name}",
	query: new RecommendInput { Positive = { new float[] { 0.1f, 0.2f, 0.3f } }, Negative = { 0 } }
);

// Fusion query
await client.QueryAsync(
	collectionName: "{collection_name}",
	prefetch: new List<PrefetchQuery>
	{
		new()
		{
			Query = new (float, uint)[] { (0.22f, 1), (0.8f, 42), },
			Using = "sparse",
			Limit = 20
		},
		new()
		{
			Query = new float[] { 0.01f, 0.45f, 0.67f },
			Using = "dense",
			Limit = 20
		}
	},
	query: Fusion.Rrf
);

// 2-stage query
await client.QueryAsync(
	collectionName: "{collection_name}",
	prefetch: new List<PrefetchQuery>
	{
		new() { Query = new float[] { 0.01f, 0.45f, 0.67f }, Limit = 100 }
	},
	query: new float[][] { [0.1f, 0.2f], [0.2f, 0.1f], [0.8f, 0.9f] },
	usingVector: "colbert",
	limit: 10
);
