import static io.qdrant.client.PointIdFactory.id;
import static io.qdrant.client.VectorFactory.vector;

import java.util.Arrays;
import java.util.List;

import io.qdrant.client.QdrantClient;
import io.qdrant.client.QdrantGrpcClient;

import io.qdrant.client.grpc.Points.ContextExamplePair;
import io.qdrant.client.grpc.Points.DiscoverPoints;
import io.qdrant.client.grpc.Points.TargetVector;
import io.qdrant.client.grpc.Points.VectorExample;

QdrantClient client = new QdrantClient(QdrantGrpcClient.newBuilder("localhost", 6334, false).build());

List <DiscoverPoints> discoverPoints = Arrays.asList(
    DiscoverPoints.newBuilder()
    .setCollectionName("{collection_name}")
    .setTarget(
        TargetVector.newBuilder()
        .setSingle(
            VectorExample.newBuilder()
            .setVector(vector(
                0.2 f,
                0.1 f,
                0.9 f,
                0.7 f))
            .build()))
    .addAllContext(
        List.of(
            ContextExamplePair.newBuilder()
            .setPositive(VectorExample
                .newBuilder()
                .setId(id(100)))
            .setNegative(VectorExample
                .newBuilder()
                .setId(id(718)))
            .build(),
            ContextExamplePair.newBuilder()
            .setPositive(VectorExample
                .newBuilder()
                .setId(id(200)))
            .setNegative(VectorExample
                .newBuilder()
                .setId(id(300)))
            .build()))
    .setLimit(10)
    .build(),
    DiscoverPoints.newBuilder()
    .setCollectionName("{collection_name}")
    .setTarget(
        TargetVector.newBuilder()
        .setSingle(
            VectorExample.newBuilder()
            .setVector(vector(
                0.5 f, 0.3 f, 0.2 f, 0.3 f))
            .build()))
    .addAllContext(
        List.of(
            ContextExamplePair.newBuilder()
            .setPositive(VectorExample
                .newBuilder()
                .setId(id(342)))
            .setNegative(VectorExample
                .newBuilder()
                .setId(id(213)))
            .build(),
            ContextExamplePair.newBuilder()
            .setPositive(VectorExample
                .newBuilder()
                .setId(id(100)))
            .setNegative(VectorExample
                .newBuilder()
                .setId(id(200)))
            .build()))
    .setLimit(10)
    .build());
client.discoverBatchAsync("{collection_name}", discoverPoints, null);
