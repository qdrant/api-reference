import static io.qdrant.client.QueryFactory.nearest;

import java.util.List;

import io.qdrant.client.QdrantClient;
import io.qdrant.client.QdrantGrpcClient;
import io.qdrant.client.grpc.Points.QueryPoints;

QdrantClient client =
    new QdrantClient(QdrantGrpcClient.newBuilder("localhost", 6334, false).build());

client
    .queryBatchAsync(
        "{collection_name}",
        List.of(
            QueryPoints.newBuilder().setQuery(nearest(0.1f, 0.2f, 0.3f)).build(),
            QueryPoints.newBuilder().setQuery(nearest(0.4f, 0.5f, 0.6f)).build()))
    .get();
