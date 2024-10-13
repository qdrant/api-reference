import static io.qdrant.client.ConditionFactory.matchKeyword;

import io.qdrant.client.QdrantClient;
import io.qdrant.client.QdrantGrpcClient;
import io.qdrant.client.grpc.Points.Filter;
import io.qdrant.client.grpc.Points.SearchMatrixPoints;

QdrantClient client =
    new QdrantClient(QdrantGrpcClient.newBuilder("localhost", 6334, false).build());

client
    .searchMatrixPairsAsync(
        SearchMatrixPoints.newBuilder()
            .setCollectionName("{collection_name}")
            .setFilter(
                Filter.newBuilder()
                    .addMust(matchKeyword("color", "red"))
                    .build()
            )
            .setSample(100)
            .setLimit(5)
            .build())
    .get();
