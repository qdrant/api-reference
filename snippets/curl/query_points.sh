# Query nearest by ID
curl -X POST \
  'http://localhost:6333/collections/collection_name/points/query' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "query": "43cf51e2-8777-4f52-bc74-c2cbde0c8b04"
}'

# Recommend on the average of these vectors
curl -X POST \
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
curl -X POST \
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
curl -X POST \
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
curl -X POST \
  'http://localhost:6333/collections/collection_name/points/query' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "query": {
    "sample": "random"
  }
}'

# Score boost depending on payload conditions (as of 1.14.0)
curl -X POST \
  'http://localhost:6333/collections/collection_name/points/query' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
    "prefetch": {
        "query": [0.2, 0.8, ...],  // <-- dense vector
        "limit": 50
    }
    "query": {
        "formula": {
            "sum": [
                "$score,
                { "mult": [ 0.5, { "key": "tag", "match": { "any": ["h1", "h2", "h3", "h4"] } } ] },
                { "mult": [ 0.25, { "key": "tag", "match": { "any": ["p", "li"] } } ] }
            ]
        }
    }
}'

# Score boost geographically closer points (as of 1.14.0)
curl -X POST \
  'http://localhost:6333/collections/collection_name/points/query' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
    "prefetch": { "query": [0.2, 0.8, ...], "limit": 50 },
    "query": {
        "formula": {
            "sum": [
                "$score",
                {
                    "gauss_decay": {
                        "x": {
                            "geo_distance": {
                                "origin": { "lat": 52.504043, "lon": 13.393236 } // Berlin
                                "to": "geo.location"
                            }
                        },
                        "scale": 5000 // 5km
                    }
                }
            ]
        },
        "defaults": { "geo.location": {"lat": 48.137154, "lon": 11.576124} } // Munich
    }
}'
