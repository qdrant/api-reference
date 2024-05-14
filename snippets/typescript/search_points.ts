import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

client.search("{collection_name}", {
    filter: {
        must: [
            {
                key: "city",
                match: {
                    value: "London",
                },
            },
        ],
    },
    vector: [0.2, 0.1, 0.9, 0.7],
    limit: 3,
});
