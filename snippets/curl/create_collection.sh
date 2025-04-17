# Create a collection with default dense vector
curl  -X PUT \
  'http://localhost:6333/collections/collection_name' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "vectors": {
    "size": 384,
    "distance": "Cosine"
  }
}'

# Create a collection with named dense and sparse vectors
curl  -X PUT \
  'http://localhost:6333/collections/collection_name' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "vectors": {
    "dense-vector-name": {
      "size": 1536,
      "distance": "Cosine"
    }
  },
  "sparse_vectors": {
    "sparse-vector-name": {
      "index": {
        "on_disk": true
      }
    }
  }
}'
