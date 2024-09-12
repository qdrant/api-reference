curl  -X PATCH \
  'http://localhost:6333/collections/collection_name' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "optimizers_config": {
    "indexing_threshold": 10000
  }
}'