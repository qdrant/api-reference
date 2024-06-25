import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

client.setPayload("global_patient_data", {
  payload: {
    is_smoker: false,
    gender: "Male",
  },
  points: [0, 3, 10],
});
