import java.util.List;

import static io.qdrant.client.QueryFactory.nearest;

import io.qdrant.client.grpc.Points.QueryPointGroups;
import io.qdrant.client.grpc.Points.QueryPoints;


QdrantClient client =
    new QdrantClient(QdrantGrpcClient.newBuilder("localhost", 6334, false).build());

client
    .queryGroupsAsync(
        QueryPointGroups.newBuilder()
            .setCollectionName("{collection_name}")
            .setGroupBy("document_id")
            .setGroupSize(5)
            .setLimit(10)
            .setQuery(nearest(List.of(0.01f, 0.45f, 0.67f)))
            .build())
    .get();
