import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

client.delete("global_patient_data", {
  points: [0, 3, 100],
});

// Delete all points with color = "red"

client.delete("global_patient_data", {
  filter: {
    must: [
      {
        key: "country",
        match: {
          value: "India",
        },
      },
    ],
  },
});
