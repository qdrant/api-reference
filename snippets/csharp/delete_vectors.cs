using Qdrant.Client;

var client = new QdrantClient("localhost", 6334);

await client.DeleteVectorsAsync("{collection_name}", ["text", "image"], [0, 3, 10]);
