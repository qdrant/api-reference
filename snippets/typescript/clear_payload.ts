import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

client.clearPayload("{collection_name}", {
  points: [0, 3, 100],
});
