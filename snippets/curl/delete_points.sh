# Delete points by IDs
curl  -X POST \
  'http://localhost:6333/collections/collection_name/points/delete' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "points": [
    0,
    3,
    100
  ]
}'

# Delete points satisfying a filter condition
curl  -X POST \
  'http://localhost:6333/collections/collection_name/points/delete' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "filter": {
    "must": [
      {
        "key": "color",
        "match": {
          "value": "red"
        }
      }
    ]
  }
}'
