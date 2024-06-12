import static io.qdrant.client.ConditionFactory.matchKeyword;

import io.qdrant.client.QdrantClient;
import io.qdrant.client.QdrantGrpcClient;

import io.qdrant.client.grpc.Points.Filter;

QdrantClient client = new QdrantClient(QdrantGrpcClient.newBuilder("localhost", 6334, false).build());

client.collectionExistsAsync("{collection_name}").get();
