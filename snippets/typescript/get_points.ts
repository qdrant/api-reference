import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

client.retrieve("global_patient_data", {
  ids: [0, 3, 100],
});
