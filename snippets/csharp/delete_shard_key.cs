using Qdrant.Client;
using Qdrant.Client.Grpc;

var client = new QdrantClient("localhost", 6334);

await client.DeleteShardKeyAsync(
    "{collection_name}",
    new DeleteShardKey { ShardKey = new ShardKey { Keyword = "shard_key", } }
);
