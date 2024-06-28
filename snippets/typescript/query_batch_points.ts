import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

client.queryBatch("{collection_name", {
    searches: [{
        query: [0.01, 0.45, 0.67]
    },
    {
        query: [0.11, 0.35, 0.6]
    }]
});
