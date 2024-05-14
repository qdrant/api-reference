import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

client.overwritePayload("{collection_name}", {
  payload: {
    property1: "string",
    property2: "string",
  },
  points: [0, 3, 10],
});
