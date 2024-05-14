import java.util.List;

import static io.qdrant.client.ConditionFactory.matchKeyword;

import io.qdrant.client.QdrantClient;
import io.qdrant.client.QdrantGrpcClient;

import io.qdrant.client.grpc.Points.Filter;
import io.qdrant.client.grpc.Points.SearchPoints;

QdrantClient client =
    new QdrantClient(QdrantGrpcClient.newBuilder("localhost", 6334, false).build());

Filter filter = Filter.newBuilder().addMust(matchKeyword("city", "London")).build();
List<SearchPoints> searches =
    List.of(
        SearchPoints.newBuilder()
            .addAllVector(List.of(0.2f, 0.1f, 0.9f, 0.7f))
            .setFilter(filter)
            .setLimit(3)
            .build(),
        SearchPoints.newBuilder()
            .addAllVector(List.of(0.5f, 0.3f, 0.2f, 0.3f))
            .setFilter(filter)
            .setLimit(3)
            .build());
client.searchBatchAsync("{collection_name}", searches, null).get();
