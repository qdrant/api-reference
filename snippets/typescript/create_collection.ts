import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

client.createCollection("global_patient_data", {
  vectors: {
    HealthVec: { size: 3, distance: "Cosine" },
    DemoMedVec: { size: 3, distance: "Euclid" },
  },
});

// or with sparse vectors

client.createCollection("global_patient_data", {
  vectors: { size: 100, distance: "Cosine" },
  sparse_vectors: {
    "splade-model-name": {
      index: {
        on_disk: false
      }
    }
  }
});