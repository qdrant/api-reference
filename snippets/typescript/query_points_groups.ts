import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

client.queryGroups("{collection_name}", {
    query: [0.01, 0.45, 0.67],
    group_by: "document_id",
    limit: 10,
    group_size: 5,
});
