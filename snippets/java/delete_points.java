import static io.qdrant.client.PointIdFactory.id;
import static io.qdrant.client.ConditionFactory.matchKeyword;

import java.util.List;

import io.qdrant.client.QdrantClient;
import io.qdrant.client.QdrantGrpcClient;

import io.qdrant.client.grpc.Common.Filter;

QdrantClient client = new QdrantClient(
                QdrantGrpcClient.newBuilder("localhost", 6334, false).build());

client.deleteAsync("{collection_name}", List.of(id(0), id(3), id(100)));

client
    .deleteAsync(
        "{collection_name}",
        Filter.newBuilder().addMust(matchKeyword("color", "red")).build())
    .get();
