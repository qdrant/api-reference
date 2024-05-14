import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

client.searchPointGroups("{collection_name}", {
    vector: [1.1],
    group_by: "document_id",
    limit: 4,
    group_size: 2,
});
