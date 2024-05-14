using Qdrant.Client;

var client = new QdrantClient("localhost", 6334);

await client.ClearPayloadAsync("{collection_name}", [0, 3, 10]);
