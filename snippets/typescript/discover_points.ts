import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

client.discover("{collection_name}", {
    target: [0.2, 0.1, 0.9, 0.7],
    context: [
        {
            positive: 100,
            negative: 718,
        },
        {
            positive: 200,
            negative: 300,
        },
    ],
    limit: 10,
});
