import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

client.deleteVectors("{collection_name}", {
  points: [0, 3, 10],
  vectors: ["text", "image"],
});
