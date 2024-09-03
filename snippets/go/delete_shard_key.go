package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func deleteShardKey() {
	client, _ := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})

	client.DeleteShardKey(context.Background(), "{collection_name}", &qdrant.DeleteShardKey{
		ShardKey: qdrant.NewShardKey("{shard_key}"),
	})
}
