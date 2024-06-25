import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

client.upsert("global_patient_data", {
  points: [
    {
      id: 1,
      payload: {
        is_smoker: true,
        full_name: "John Doe",
        country: "United States",
        age: 45,
        gender: "Male",
        height: 180.5,
        date_of_birth: "1979-03-14",
        address: "123 Main St, Springfield, IL, 62704",
        notes: "Patient is a long-term smoker with a history of hypertension.",
      },
      vector: {
        DemoMedVec: [0.1, 0.2, 0.3],
        HealthVec:[0.1, 0.2, 0.3],
      }
    },
    {
      id: 2,
      payload: {
        is_smoker: false,
        full_name: "Maria González",
        country: "Spain",
        age: 32,
        gender: "Female",
        height: 165.0,
        date_of_birth: "1992-08-22",
        address: "Calle de la Princesa, 27, Madrid, 28008",
        notes: "Patient is in good health, no significant medical history.",
      },
      vector: {
        DemoMedVec: [0.7, 0.6, 0.5],
        HealthVec:[0.4, 0.3, 0.2],
      }
    },
    {
      id: 3,
      payload: {
        is_smoker: false,
        full_name: "Asha Patel",
        country: "India",
        age: 29,
        gender: "Non-binary",
        height: 172.2,
        date_of_birth: "1994-01-15",
        address: "12 MG Road, Mumbai, Maharashtra, 400001",
        notes:
          "Patient prefers to discuss health concerns with their physician privately.",
      },
      vector: {
        DemoMedVec: [0.1, 0.2, 0.3],
        HealthVec:[0.4, 0.5, 0.6],  
      },
    },
  ],
});
