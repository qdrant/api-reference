import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

client.scroll("{collection_name}", {
    filter: {
        must: [
            {
                key: "country",
                match: {
                    value: "Canada",
                },
            },
        ],
    },
    limit: 1,
    with_payload: true,
    with_vector: false,
});
