import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

// Query nearest by ID
client.query("{collection_name", {
    query: "43cf51e2-8777-4f52-bc74-c2cbde0c8b04"
});

// Recommend on the average of these vectors
client.query("{collection_name}", {
    query: {
        recommend: {
            positive: ["43cf51e2-8777-4f52-bc74-c2cbde0c8b04", [0.11, 0.35, 0.6]],
            negative: [0.01, 0.45, 0.67]
        }
    }
});

// Fusion query
client.query("{collection_name}", {
    prefetch: [
        {
            query: {
                values: [0.22, 0.8],
                indices: [1, 42],
            },
            using: 'sparse',
            limit: 20,
        },
        {
            query: [0.01, 0.45, 0.67],
            using: 'dense',
            limit: 20,
        },
    ],
    query: {
        fusion: 'rrf',
    },
});

// 2-stage query
client.query("{collection_name}", {
    prefetch: {
        query: [1, 23, 45, 67],
        limit: 100,
    },
    query: [
        [0.1, 0.2],
        [0.2, 0.1],
        [0.8, 0.9],
    ],
    using: 'colbert',
    limit: 10,
});
