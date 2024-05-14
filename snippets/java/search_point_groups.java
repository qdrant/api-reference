import java.util.List;

import io.qdrant.client.QdrantClient;
import io.qdrant.client.QdrantGrpcClient;

import io.qdrant.client.grpc.Points.SearchPointGroups;

QdrantClient client = new QdrantClient(
                QdrantGrpcClient.newBuilder("localhost", 6334, false).build());

client
    .searchGroupsAsync(
        SearchPointGroups.newBuilder()
            .setCollectionName("{collection_name}")
            .addAllVector(List.of(1.1f))
            .setGroupBy("document_id")
            .setLimit(4)
            .setGroupSize(2)
            .build())
    .get();
