using Qdrant.Client;

var client = new QdrantClient("localhost", 6334);

await client.CreateAliasAsync(aliasName: "production_collection", collectionName: "example_collection");

await client.DeleteAliasAsync("production_collection");
