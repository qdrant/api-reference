import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

client.recommendPointGroups("{collection_name}", {
    positive: [100, 231],
    negative: [718],
    group_by: "document_id",
    limit: 3,
    group_size: 2,
});
