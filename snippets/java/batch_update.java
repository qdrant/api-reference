import java.util.List;
import java.util.Map;

import static io.qdrant.client.PointIdFactory.id;
import static io.qdrant.client.ValueFactory.value;
import static io.qdrant.client.VectorsFactory.vectors;

import io.qdrant.client.QdrantClient;
import io.qdrant.client.QdrantGrpcClient;

import io.qdrant.client.grpc.Points.PointStruct;
import io.qdrant.client.grpc.Points.PointVectors;
import io.qdrant.client.grpc.Points.PointsIdsList;
import io.qdrant.client.grpc.Points.PointsSelector;
import io.qdrant.client.grpc.Points.PointsUpdateOperation;
import io.qdrant.client.grpc.Points.PointsUpdateOperation.PointStructList;
import io.qdrant.client.grpc.Points.PointsUpdateOperation.SetPayload;
import io.qdrant.client.grpc.Points.PointsUpdateOperation.UpdateVectors;

QdrantClient client = new QdrantClient(
                QdrantGrpcClient.newBuilder("localhost", 6334, false).build());

client
    .batchUpdateAsync(
        "{collection_name}",
        List.of(
            PointsUpdateOperation.newBuilder()
            .setUpsert(
                PointStructList.newBuilder()
                .addPoints(
                    PointStruct.newBuilder()
                    .setId(id(1))
                    .setVectors(vectors(
                        1.0 f,
                        2.0 f,
                        3.0 f,
                        4.0 f))
                    .build())
                .build())
            .build(),
            PointsUpdateOperation.newBuilder()
            .setUpdateVectors(
                UpdateVectors.newBuilder()
                .addPoints(
                    PointVectors.newBuilder()
                    .setId(id(1))
                    .setVectors(vectors(
                        1.0 f,
                        2.0 f,
                        3.0 f,
                        4.0 f))
                    .build())
                .build())
            .build(),
            PointsUpdateOperation.newBuilder()
            .setSetPayload(
                SetPayload.newBuilder()
                .setPointsSelector(
                    PointsSelector.newBuilder()
                    .setPoints(PointsIdsList
                        .newBuilder()
                        .addIds(id(1))
                        .build())
                    .build())
                .putAllPayload(
                    Map.of("test_payload_2",
                        value(2),
                        "test_payload_3",
                        value(3)))
                .build())
            .build()))
    .get();
