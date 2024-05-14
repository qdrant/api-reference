import static io.qdrant.client.PointIdFactory.id;

import java.util.List;

import io.qdrant.client.QdrantClient;
import io.qdrant.client.QdrantGrpcClient;

import io.qdrant.client.grpc.Points.RecommendPointGroups;
import io.qdrant.client.grpc.Points.RecommendStrategy;

QdrantClient client =
    new QdrantClient(QdrantGrpcClient.newBuilder("localhost", 6334, false).build());

client.recommendGroupsAsync(RecommendPointGroups.newBuilder()
                .setCollectionName("{collection_name}")
                .setGroupBy("document_id")
                .setGroupSize(2)
                .addAllPositive(List.of(id(100), id(200)))
                .addAllNegative(List.of(id(718)))
                .setStrategy(RecommendStrategy.AverageVector)
                .setLimit(3)
                .build());
