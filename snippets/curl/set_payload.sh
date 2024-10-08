# Set payload by ID
curl  -X POST \
  'http://localhost:6333/collections/collection_name/points/payload' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "payload": {
    "property1": "some_value",
    "property2": 32,
    "property3": true
  },
  "points": [
    0,
    3,
    10
  ]
}'

# Set payload by filter
curl  -X POST \
  'http://localhost:6333/collections/collection_name/points/payload' \
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
  "payload": {
    "property1": "some_value",
    "property2": 32,
    "property3": true
  }
}'