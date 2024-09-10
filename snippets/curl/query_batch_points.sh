curl  -X POST \
  'http://localhost:6333/collections/collection_name/points/query/batch' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "searches": [
    {
      "query": [
        0.23,
        0.325,
        0.623
      ],
      "limit": 1
    },
    {
      "query": [
        0.423,
        0.23,
        0.623
      ],
      "limit": 5,
      "using": "image-vector"
    }
  ]
}'