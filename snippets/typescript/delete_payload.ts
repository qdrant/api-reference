import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

client.deletePayload("{collection_name}", {
  keys: ["color", "price"],
  points: [0, 3, 100],
});
