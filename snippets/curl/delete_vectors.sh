# Delete vectors by ID
curl  -X POST \
  'http://localhost:6333/collections/collection_name/points/vectors/delete' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "points": [
    0,
    3,
    10
  ],
  "vectors": [
    "text",
    "image"
  ]
}'

# Delete vectors by filter
curl  -X POST \
  'http://localhost:6333/collections/collection_name/points/vectors/delete' \
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
  "vectors": [
    "text",
    "image"
  ]
}'