curl  -X POST \
  'http://localhost:6333/collections/collection_name/points/scroll' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "with_payload": [
    "city",
    "color"
  ],
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
  "limit": 2,
  "with_vector": false
}'