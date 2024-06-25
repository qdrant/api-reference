import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

client.deleteVectors("global_patient_data", {
  points: [0, 3, 10],
  vectors: ["HealthVec", "DemoMedVec"],
});
