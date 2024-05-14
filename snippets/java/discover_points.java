import java.util.List;

import static io.qdrant.client.PointIdFactory.id;
import static io.qdrant.client.VectorFactory.vector;

import io.qdrant.client.QdrantClient;
import io.qdrant.client.QdrantGrpcClient;

import io.qdrant.client.grpc.Points.ContextExamplePair;
import io.qdrant.client.grpc.Points.DiscoverPoints;
import io.qdrant.client.grpc.Points.TargetVector;
import io.qdrant.client.grpc.Points.VectorExample;

QdrantClient client =
    new QdrantClient(QdrantGrpcClient.newBuilder("localhost", 6334, false).build());

client
    .discoverAsync(
        DiscoverPoints.newBuilder()
            .setCollectionName("{collection_name}")
            .setTarget(
                TargetVector.newBuilder()
                    .setSingle(
                        VectorExample.newBuilder()
                            .setVector(vector(0.2f, 0.1f, 0.9f, 0.7f))
                            .build()))
            .addAllContext(
                List.of(
                    ContextExamplePair.newBuilder()
                        .setPositive(VectorExample.newBuilder().setId(id(100)))
                        .setNegative(VectorExample.newBuilder().setId(id(718)))
                        .build(),
                    ContextExamplePair.newBuilder()
                        .setPositive(VectorExample.newBuilder().setId(id(200)))
                        .setNegative(VectorExample.newBuilder().setId(id(300)))
                        .build()))
            .setLimit(10)
            .build())
    .get();
