import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

client.count("global_patient_data", {
    filter: {
        must: [
            {
                key: "is_smoker",
                match: {
                    value: true,
                },
            },
        ],
    },
    exact: true,
});
