import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

client.createCollection("{collection_name}", {
  vectors: { size: 100, distance: "Cosine" },
});

// or with sparse vectors

client.createCollection("{collection_name}", {
  vectors: { size: 100, distance: "Cosine" },
  sparse_vectors: {
    "splade-model-name": {
      index: {
        on_disk: false
      }
    }
  }
});