using Qdrant.Client;

var client = new QdrantClient("localhost", 6334);

await client.QueryGroupsAsync(
	collectionName: "{collection_name}",
    groupBy: "document_id",
    groupSize: 5,
    limit: 10,
	query: new float[] { 0.01f, 0.45f, 0.67f }
);
