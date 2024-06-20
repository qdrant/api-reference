using Qdrant.Client;

var client = new QdrantClient("localhost", 6334);

await client.DeleteAsync(collectionName: "{collection_name}", ids: [0, 3, 100]);

// Delete all points with color = "red"

await client.DeleteAsync(collectionName: "{collection_name}", filter: MatchKeyword("color", "red"));
