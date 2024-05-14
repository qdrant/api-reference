import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

const searches = [
    {
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
    },
    {
        target: [0.5, 0.3, 0.2, 0.3],
        context: [
            {
                positive: 342,
                negative: 213,
            },
            {
                positive: 100,
                negative: 200,
            },
        ],
        limit: 5,
    },
];

client.discoverBatchPoints("{collection_name}", {
    searches,
});
