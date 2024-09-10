curl  -X POST \
  'http://localhost:6333/collections/collection_name/points/search/batch' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "searches": [
    {
      "vector": [
        0.2,
        0.1,
        0.9,
        0.7
      ],
      "limit": 3
    },
    {
      "vector": [
        0.5,
        0.3,
        0.2,
        0.3
      ],
      "filter": {
        "must": [
          {
            "key": "city",
            "match": {
              "value": "London"
            }
          }
        ]
      },
      "limit": 3
    }
  ]
}'