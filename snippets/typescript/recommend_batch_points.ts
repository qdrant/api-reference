import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

const searches = [
    {
        positive: [100, 231],
        negative: [718],
        limit: 3,
    },
    {
        positive: [200, 67],
        negative: [300],
        limit: 3,
    },
];

client.recommend_batch("{collection_name}", {
    searches,
});
