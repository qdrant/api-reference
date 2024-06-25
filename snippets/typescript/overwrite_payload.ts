import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

client.overwritePayload("global_patient_data", {
  payload: {
    is_smoker: false,
    full_name: "Hiroshi Tanaka",
    country: "Japan",
    age: 58,
    gender: "Male",
    height: 175.3,
    date_of_birth: "1966-05-21",
    address: "1-2-3 Shibuya, Tokyo, 150-0002",
    notes:
      "Patient has a family history of diabetes. Regular check-ups recommended.",
  },
  points: [0],
});
