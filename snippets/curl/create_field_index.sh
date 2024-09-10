curl  -X PUT \
  'http://localhost:6333/collections/collection_name/index' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "field_name": "field_name",
  "field_schema": "keyword"
}'