import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

client.retrieve("{collection_name}", {
  ids: [0, 3, 100],
});
