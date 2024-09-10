curl  -X POST \
  'http://localhost:6333/collections/collection_name/points/discover/batch' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "searches": [
    {
      "target": [
        0.2,
        0.1,
        0.9,
        0.7
      ],
      "context": [
        {
          "positive": "7f6652c4-89bd-40e0-ab1d-510f7fcddd2b",
          "negative": "2c2c9f86-0171-4bd2-9ab0-b4754682cddd"
        }
      ],
      "limit": 1
    },
    {
      "target": [
        0.5,
        0.3,
        0.2,
        0.3
      ],
      "context": [
        {
          "positive": 342,
          "negative": 213
        }
      ],
      "limit": 1
    }
  ]
}'