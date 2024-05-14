import static io.qdrant.client.ShardKeyFactory.shardKey;

import io.qdrant.client.QdrantClient;
import io.qdrant.client.QdrantGrpcClient;

import io.qdrant.client.grpc.Collections.DeleteShardKey;
import io.qdrant.client.grpc.Collections.DeleteShardKeyRequest;

QdrantClient client = new QdrantClient(
                QdrantGrpcClient.newBuilder("localhost", 6334, false).build());

client.deleteShardKeyAsync(DeleteShardKeyRequest.newBuilder()
                .setCollectionName("{collection_name}")
                .setRequest(DeleteShardKey.newBuilder()
                                .setShardKey(shardKey("{shard_key}"))
                                .build())
                .build()).get();
