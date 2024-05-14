import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

client.updateCollectionAliases({
  actions: [
    {
      create_alias: {
        collection_name: "example_collection",
        alias_name: "production_collection",
      },
    },
  ],
});

client.updateCollectionAliases({
  actions: [
    {
      delete_alias: {
        alias_name: "production_collection",
      },
    },
  ],
});
