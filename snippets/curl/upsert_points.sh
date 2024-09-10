curl  -X PUT \
  'http://localhost:6333/collections/collection_name/points' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "points": [
    {
      "id": "76874cce-1fb9-4e16-9b0b-f085ac06ed6f",
      "payload": {
        "color": "red"
      },
      "vector": [
        0.9,
        0.1,
        0.1
      ]
    },
    {
      "id": 1,
      "payload": {
        "color": "green"
      },
      "vector": [
        0.1,
        0.9,
        0.1
      ]
    },
    {
      "id": 2,
      "payload": {
        "color": "blue"
      },
      "vector": [
        0.1,
        0.1,
        0.9
      ]
    }
  ]
}'