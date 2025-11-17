import static io.qdrant.client.ConditionFactory.matchKeyword;

import io.qdrant.client.QdrantClient;
import io.qdrant.client.QdrantGrpcClient;

import io.qdrant.client.grpc.Common.Filter;

QdrantClient client = new QdrantClient(QdrantGrpcClient.newBuilder("localhost", 6334, false).build());

client
    .countAsync(
        "{collection_name}",
        Filter.newBuilder().addMust(matchKeyword("color", "red")).build(),
        true)
    .get();
