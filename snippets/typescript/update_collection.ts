import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

client.updateCollection("{collection_name}", {
  optimizers_config: {
    indexing_threshold: 10000,
  },
});
