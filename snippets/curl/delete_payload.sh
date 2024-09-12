# Delete payload by ID
curl  -X POST \
  'http://localhost:6333/collections/collection_name/points/payload/delete' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "keys": [
    "color",
    "price"
  ],
  "points": [
    0,
    3,
    100
  ]
}'

# Delete payload by filter
curl  -X POST \
  'http://localhost:6333/collections/collection_name/points/payload/delete' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "keys": [
    "color",
    "price"
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
  }
}'