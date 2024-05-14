import static io.qdrant.client.ShardKeyFactory.shardKey;

import io.qdrant.client.QdrantClient;
import io.qdrant.client.QdrantGrpcClient;

import io.qdrant.client.grpc.Collections.DeleteShardKey;
import io.qdrant.client.grpc.Collections.DeleteShardKeyRequest;

QdrantClient client = new QdrantClient(
                QdrantGrpcClient.newBuilder("localhost", 6334, false).build());

client.createAliasAsync("production_collection", "example_collection").get();

client.deleteAliasAsync("production_collection").get();
