import java.util.List;

import static io.qdrant.client.ConditionFactory.matchKeyword;
import static io.qdrant.client.PointIdFactory.id;

import io.qdrant.client.QdrantClient;
import io.qdrant.client.QdrantGrpcClient;

import io.qdrant.client.grpc.Common.Filter;
import io.qdrant.client.grpc.Points.RecommendPoints;

QdrantClient client =
    new QdrantClient(QdrantGrpcClient.newBuilder("localhost", 6334, false).build());

Filter filter = Filter.newBuilder().addMust(matchKeyword("city", "London")).build();

List<RecommendPoints> recommendQueries =
    List.of(
        RecommendPoints.newBuilder()
            .addAllPositive(List.of(id(100), id(231)))
            .addAllNegative(List.of(id(718)))
            .setFilter(filter)
            .setLimit(3)
            .build(),
        RecommendPoints.newBuilder()
            .addAllPositive(List.of(id(200), id(67)))
            .addAllNegative(List.of(id(300)))
            .setFilter(filter)
            .setLimit(3)
            .build());

client.recommendBatchAsync("{collection_name}", recommendQueries, null).get();
