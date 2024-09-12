curl  -X POST \
  'http://localhost:6333/collections/collection_name/points' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "ids": [
    0,
    3,
    100
  ]
}'