using Qdrant.Client;

var client = new QdrantClient("localhost", 6334);

await client.SearchGroupsAsync(
  collectionName: "{collection_name}",
  vector: new float[] { 1.1f },
  groupBy: "document_id",
  limit: 4,
  groupSize: 2
);
