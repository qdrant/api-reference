using Qdrant.Client;

var client = new QdrantClient("localhost", 6334);

await client.ListFullSnapshotsAsync();
