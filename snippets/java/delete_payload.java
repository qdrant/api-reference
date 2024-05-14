import java.util.List;

import static io.qdrant.client.PointIdFactory.id;

import io.qdrant.client.QdrantClient;
import io.qdrant.client.QdrantGrpcClient;

QdrantClient client = new QdrantClient(
                QdrantGrpcClient.newBuilder("localhost", 6334, false).build());

client
    .deletePayloadAsync(
        "{collection_name}",
        List.of("color", "price"),
        List.of(id(0), id(3), id(100)),
        true,
        null,
        null)
    .get();
