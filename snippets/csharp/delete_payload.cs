using Qdrant.Client;

var client = new QdrantClient("localhost", 6334);

await client.DeletePayloadAsync("{collection_name}", ["color", "price"], [0, 3, 10]);
