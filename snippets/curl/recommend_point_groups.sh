curl  -X POST \
  'http://localhost:6333/collections/collection_name/points/recommend/groups' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "positive": [
    100,
    231
  ],
  "negative": [
    718
  ],
  "group_by": "document_id",
  "limit": 3,
  "group_size": 2
}'