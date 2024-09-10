curl  -X PUT \
  'http://localhost:6333/collections/collection_name/points/vectors?wait=true' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "points": [
    {
      "id": 1,
      "vector": [
        0.0984,
        0.1406,
        0.8973
      ]
    },
    {
      "id": 2,
      "vector": {
        "vector-name": [
          0.9,
          0.8,
          0.7,
          0.6
        ]
      }
    }
  ]
}'