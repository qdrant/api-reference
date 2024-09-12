curl  -X POST \
  'http://localhost:6333/collections/collection_name/points/batch' \
  --header 'Content-Type: application/json' \
  --header 'api-key: <api-key-value>' \
  --data-raw '{
  "operations": [
    {
      "upsert": {
        "points": [
          {
            "id": 1,
            "vector": [
              0.4,
              0.3,
              0.2,
              0.1
            ]
          }
        ]
      }
    },
    {
      "update_vectors": {
        "points": [
          {
            "id": 1,
            "vector": [
              0.11,
              0.22,
              0.33,
              0.44
            ]
          }
        ]
      }
    },
    {
      "set_payload": {
        "payload": {
          "test_payload_2": 2,
          "test_payload_3": 3
        },
        "points": [
          1
        ]
      }
    }
  ]
}'