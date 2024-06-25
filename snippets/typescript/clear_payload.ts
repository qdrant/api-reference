import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

client.clearPayload("global_patient_data", {
  points: [0, 3, 100],
});
