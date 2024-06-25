import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

client.updateVectors("global_patient_data", {
  points: [
    {
      id: 1,
      vector: {
        HealthVec: [0.1, 0.2, 0.3],
      },
    },
    {
      id: 2,
      vector: {
        DemoMedVec: [0.9, 0.8, 0.7],
      },
    },
  ],
});
