import static io.qdrant.client.QueryFactory.fusion;
import static io.qdrant.client.QueryFactory.nearest;
import static io.qdrant.client.QueryFactory.recommend;
import static io.qdrant.client.VectorInputFactory.vectorInput;

import java.util.UUID;

import io.qdrant.client.QdrantClient;
import io.qdrant.client.QdrantGrpcClient;
import io.qdrant.client.grpc.Points.Fusion;
import io.qdrant.client.grpc.Points.PrefetchQuery;
import io.qdrant.client.grpc.Points.QueryPoints;
import io.qdrant.client.grpc.Points.RecommendInput;

QdrantClient client =
    new QdrantClient(QdrantGrpcClient.newBuilder("localhost", 6334, false).build());

// Query nearest by ID
client
    .queryAsync(
        QueryPoints.newBuilder()
            .setCollectionName("{collection_name}")
            .setQuery(nearest(UUID.fromString("43cf51e2-8777-4f52-bc74-c2cbde0c8b04")))
            .build())
    .get();

// Recommend on the average of these vectors
client
    .queryAsync(
        QueryPoints.newBuilder()
            .setCollectionName("{collection_name}")
            .setQuery(
                recommend(
                    RecommendInput.newBuilder()
                        .addPositive(vectorInput(UUID.fromString("43cf51e2-8777-4f52-bc74-c2cbde0c8b04")))
                        .addPositive(vectorInput(0.11f, 0.35f, 0.6f))
                        .addNegative(vectorInput(0.01f, 0.45f, 0.67f))
                        .build()))
            .build())
    .get();

// Fusion query
client
    .queryAsync(
        QueryPoints.newBuilder()
            .setCollectionName("{collection_name}")
            .addPrefetch(
                PrefetchQuery.newBuilder()
                    .setQuery(nearest(List.of(0.22f, 0.8f), List.of(1, 42)))
                    .setUsing("sparse")
                    .setLimit(20)
                    .build())
            .addPrefetch(
                PrefetchQuery.newBuilder()
                    .setQuery(nearest(List.of(0.01f, 0.45f, 0.67f)))
                    .setUsing("dense")
                    .setLimit(20)
                    .build())
            .setQuery(fusion(Fusion.RRF))
            .build())
    .get();

// 2-stage query
client
    .queryAsync(
        QueryPoints.newBuilder()
            .setCollectionName("{collection_name}")
            .addPrefetch(
                PrefetchQuery.newBuilder()
                    .setQuery(nearest(0.01f, 0.45f, 0.67f))
                    .setLimit(100)
                    .build())
            .setQuery(
                nearest(
                    new float[][] {
                      {0.1f, 0.2f},
                      {0.2f, 0.1f},
                      {0.8f, 0.9f}
                    }))
            .setUsing("colbert")
            .setLimit(10)
            .build())
    .get();

// Random sampling (as of 1.11.0)
client
    .queryAsync(
        QueryPoints.newBuilder()
            .setCollectionName("{collection_name}")
            .setQuery(sample(Sample.Random))
            .build())
    .get();
