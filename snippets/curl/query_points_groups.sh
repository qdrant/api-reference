curl  -X POST \
  'http://localhost:6333/collections/collection_name/points/query/groups' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "query": [
    0.324,
    0.643,
    0.423
  ],
  "group_by": "document_id",
  "limit": 1,
  "group_size": 5
}'