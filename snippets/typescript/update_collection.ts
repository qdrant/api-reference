import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

client.updateCollection("global_patient_data", {
  optimizers_config: {
    indexing_threshold: 10000,
  },
});
