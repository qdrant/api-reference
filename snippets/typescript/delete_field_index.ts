import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

client.deletePayloadIndex("{collection_name}", "{field_name}");
