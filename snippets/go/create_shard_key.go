package client

import (
	"context"

	"github.com/qdrant/go-client/qdrant"
)

func createShardKey() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "localhost",
		Port: 6334,
	})
	if err != nil {
		panic(err)
	}

	err = client.CreateShardKey(context.Background(), "{collection_name}", &qdrant.CreateShardKey{
		ShardKey: qdrant.NewShardKey("{shard_key}"),
	})
	if err != nil {
		panic(err)
	}
}
