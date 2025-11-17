import static io.qdrant.client.ConditionFactory.matchKeyword;

import io.qdrant.client.QdrantClient;
import io.qdrant.client.QdrantGrpcClient;

import io.qdrant.client.grpc.Common.Filter;
import io.qdrant.client.grpc.Points.SearchPoints;

QdrantClient client = new QdrantClient(
                QdrantGrpcClient.newBuilder("localhost", 6334, false).build());

client
    .searchAsync(
        SearchPoints.newBuilder()
            .setCollectionName("{collection_name}")
            .setFilter(Filter.newBuilder().addMust(matchKeyword("city", "London")).build())
            .addAllVector(List.of(0.2f, 0.1f, 0.9f, 0.7f))
            .setLimit(3)
            .build())
    .get();
