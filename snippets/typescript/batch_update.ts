import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

client.batchUpdate("{collection_name}", {
    operations: [
        {
            upsert: {
                points: [
                    {
                        id: 1,
                        vector: [1.0, 2.0, 3.0, 4.0],
                        payload: {},
                    },
                ],
            },
        },
        {
            update_vectors: {
                points: [
                    {
                        id: 1,
                        vector: [1.0, 2.0, 3.0, 4.0],
                    },
                ],
            },
        },
        {
            set_payload: {
                payload: {
                    test_payload_2: 2,
                    test_payload_3: 3,
                },
                points: [1],
            },
        },
    ],
});
