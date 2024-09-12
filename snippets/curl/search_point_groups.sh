curl  -X POST \
  'http://localhost:6333/collections/collection_name/points/search/groups' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "vector": [
    0.2,
    0.1,
    0.9,
    0.7
  ],
  "group_by": "document_id",
  "limit": 4,
  "group_size": 2
}'