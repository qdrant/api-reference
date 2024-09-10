curl  -X PUT \
  'http://localhost:6333/collections/collection_name/shards' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "shard_key": "{shard_key}"
}'