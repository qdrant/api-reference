package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func createShardKey() {
	client, _ := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})

	client.CreateShardKey(context.Background(), "{collection_name}", &qdrant.CreateShardKey{
		ShardKey: qdrant.NewShardKey("{shard_key}"),
	})
}
