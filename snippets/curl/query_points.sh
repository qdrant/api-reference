# Query nearest by ID
curl  -X POST \
  'http://localhost:6333/collections/collection_name/points/query' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "query": "43cf51e2-8777-4f52-bc74-c2cbde0c8b04"
}'

# Recommend on the average of these vectors
curl  -X POST \
  'http://localhost:6333/collections/collection_name/points/query' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "query": {
    "recommend": {
      "positive": [
        [
          0.11,
          0.35,
          0.6
        ]
      ],
      "negative": [
        "43cf51e2-8777-4f52-bc74-c2cbde0c8b04",
        [
          0.01,
          0.45,
          0.67
        ]
      ]
    }
  }
}'

# Fusion query
curl  -X POST \
  'http://localhost:6333/collections/collection_name/points/query' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "prefetch": [
    {
      "query": {
        "values": [
          0.22,
          0.8
        ],
        "indices": [
          1,
          42
        ]
      },
      "using": "sparse",
      "limit": 20
    },
    {
      "query": [
        0.01,
        0.45,
        0.67
      ],
      "using": "dense",
      "limit": 20
    }
  ],
  "query": {
    "fusion": "rrf"
  }
}'

# 2-stage query
curl  -X POST \
  'http://localhost:6333/collections/collection_name/points/query' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "prefetch": {
    "query": [
      1,
      23,
      45,
      67
    ],
    "limit": 100
  },
  "query": [
    [
      0.1,
      0.2,
      0.3
    ],
    [
      0.2,
      0.1,
      0.35
    ],
    [
      0.8,
      0.9,
      0.53
    ]
  ],
  "using": "colbert",
  "limit": 10
}'

# Random sampling (as of 1.11.0)
curl  -X POST \
  'http://localhost:6333/collections/collection_name/points/query' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "query": {
    "sample": "random"
  }
}'
