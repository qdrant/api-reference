import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

client.searchMatrixPairs("{collection_name}", {
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
    sample: 100,
    limit: 5,
});