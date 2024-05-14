import io.qdrant.client.QdrantClient;
import io.qdrant.client.QdrantGrpcClient;

import io.qdrant.client.grpc.Collections.PayloadSchemaType;

QdrantClient client = new QdrantClient(
                QdrantGrpcClient.newBuilder("localhost", 6334, false).build());

client.createPayloadIndexAsync(
                "{collection_name}",
                "{field_name}",
                PayloadSchemaType.Keyword,
                null,
                true,
                null,
                null);
