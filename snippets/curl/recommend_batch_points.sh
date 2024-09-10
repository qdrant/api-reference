curl  -X POST \
  'http://localhost:6333/collections/collection_name/points/recommend/batch' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "searches": [
    {
      "positive": [
        100,
        231
      ],
      "negative": [
        "9ad0884a-7bfe-43c0-b88f-c1d9a588b7e1"
      ],
      "limit": 1
    },
    {
      "positive": [
        200,
        67
      ],
      "negative": [
        300
      ],
      "limit": 5
    }
  ]
}'