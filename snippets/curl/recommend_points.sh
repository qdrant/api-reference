curl  -X POST \
  'http://localhost:6333/collections/collection_name/points/recommend' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "positive": [
    100,
    231
  ],
  "negative": [
    "3a6e2150-f7b7-496e-92cd-687e63e215fd",
    [
      0.2,
      0.3,
      0.4,
      0.5
    ]
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
  "limit": 1
}'