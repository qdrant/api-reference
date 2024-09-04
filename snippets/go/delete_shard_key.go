package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func deleteShardKey() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		panic(err)
	}

	client.DeleteShardKey(context.Background(), "{collection_name}", &qdrant.DeleteShardKey{
		ShardKey: qdrant.NewShardKey("{shard_key}"),
	})
}
