curl  -X POST \
  'http://localhost:6333/collections/collection_name/points/search' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "vector": [
    0.2,
    0.1,
    0.9,
    0.7
  ],
  "limit": 1,
  "filter": {
    "must": [
      {
        "key": "city",
        "match": {
          "value": "London"
        }
      }
    ]
  }
}'