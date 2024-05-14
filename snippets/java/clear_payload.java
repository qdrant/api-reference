import java.util.List;

import static io.qdrant.client.PointIdFactory.id;

import io.qdrant.client.QdrantClient;
import io.qdrant.client.QdrantGrpcClient;

QdrantClient client = new QdrantClient(
                QdrantGrpcClient.newBuilder("localhost", 6334, false).build());

client.clearPayloadAsync("{collection_name}", List.of(id(0), id(3), id(100)), null, null, null)
                .get();
