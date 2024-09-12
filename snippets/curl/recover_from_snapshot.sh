curl  -X PUT \
  'http://localhost:6333/collections/collection_name/snapshots/recover' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "location": "http://example.com/path/to/snapshot.shapshot"
}'