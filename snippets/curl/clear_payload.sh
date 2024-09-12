curl  -X POST \
  'http://localhost:6333/collections/collection_name/points/payload/clear' \
  --header 'Content-Type: application/json' \
  --header 'api-key: <api-key-value>' \
  --data-raw '{
  "points": [
    0,
    3,
    100
  ]
}'