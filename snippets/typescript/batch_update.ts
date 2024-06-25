import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

client.batchUpdate("global_patient_data", {
  operations: [
    {
      upsert: {
        points: [
          {
            id: 1,
            vector: {
              HealthVec: [0.1, 0.2, 0.3],
              DemoMedVec: [0.9, 0.8, 0.7],
            },
            payload: {
              is_smoker: false,
              full_name: "Emily Smith",
              country: "Canada",
              age: 38,
              gender: "Female",
              height: 165.8,
              date_of_birth: "1986-11-03",
              address: "789 Maple Avenue, Toronto, ON, M5J 2N8",
              notes: "Patient has a history of seasonal allergies.",
            },
          },
        ],
      },
    },
    {
      update_vectors: {
        points: [
          {
            id: 1,
            vector: {
                HealthVec: [0.4, 0.5, 0.6],
            },
          },
        ],
      },
    },
    {
      set_payload: {
        payload: {
            is_smoker: true,
            gender:"Male"
        },
        points: [1],
      },
    },
  ],
});
