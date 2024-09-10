# Count total number of points in a collection
curl  -X POST \
  'http://localhost:6333/collections/collection_name/points/count' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "exact": true
}'

# Count points satisfying a filter condition
curl  -X POST \
  'http://localhost:6333/collections/collection_name/points/count' \
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
  },
  "exact": true
}'