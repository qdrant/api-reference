using Qdrant.Client;

var client = new QdrantClient("localhost", 6334);

await client.DeletePayloadIndexAsync(
  collectionName: "{collection_name}",
  fieldName: "name_of_the_field_to_index"
);
