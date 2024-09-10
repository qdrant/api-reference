# Overwrite payload by ID
curl  -X PUT \
  'http://localhost:6333/collections/collection_name/points/payload' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "payload": {
    "property1": "string",
    "property2": "string"
  },
  "points": [
    0,
    3,
    10
  ]
}'

# Overwrite payload by filter
curl  -X PUT \
  'http://localhost:6333/collections/collection_name/points/payload' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "payload": {
    "property1": "string",
    "property2": "string"
  },
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