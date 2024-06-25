import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

client.deletePayload("global_patient_data", {
  keys: ["is_smoker", "notes"],
  points: [0, 3, 100],
});
