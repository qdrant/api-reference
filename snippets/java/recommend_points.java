import java.util.List;

import static io.qdrant.client.ConditionFactory.matchKeyword;
import static io.qdrant.client.PointIdFactory.id;
import static io.qdrant.client.VectorFactory.vector;

import io.qdrant.client.QdrantClient;
import io.qdrant.client.QdrantGrpcClient;

import io.qdrant.client.grpc.Common.Filter;
import io.qdrant.client.grpc.Points.RecommendPoints;
import io.qdrant.client.grpc.Points.RecommendStrategy;

QdrantClient client =
    new QdrantClient(QdrantGrpcClient.newBuilder("localhost", 6334, false).build());

client
    .recommendAsync(
        RecommendPoints.newBuilder()
            .setCollectionName("{collection_name}")
            .addAllPositive(List.of(id(100), id(200)))
            .addAllPositiveVectors(List.of(vector(100.0f, 231.0f)))
            .addAllNegative(List.of(id(718)))
            .addAllPositiveVectors(List.of(vector(0.2f, 0.3f, 0.4f, 0.5f)))
            .setStrategy(RecommendStrategy.AverageVector)
            .setFilter(Filter.newBuilder().addMust(matchKeyword("city", "London")))
            .setLimit(3)
            .build())
    .get();
