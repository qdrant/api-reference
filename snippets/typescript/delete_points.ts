import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

client.delete("{collection_name}", {
  points: [0, 3, 100],
});

// Delete all points with color = "red"

client.delete("{collection_name}", {
    filter: {
        must: [
            {
                key: "color",
                match: {
                    value: "red",
                },
            },
        ],
    },
});
