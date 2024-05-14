# Minimal curl command to create a collection with a vector field

curl -X PUT http://localhost:6333/collections/collection_name \
     -H "api-key: <apiKey>" \
     -H "Content-Type: application/json" \
     -d '{
            "vectors": {
                "size": 300,
                "distance": "Cosine"
            } 
        }'

# Or with a sparse vector field

curl -X PUT http://localhost:6333/collections/collection_name \
     -H "api-key: <apiKey>" \
     -H "Content-Type: application/json" \
     -d '{
            "vectors": {
                "size": 300,
                "distance": "Cosine"
            },
            "sparse_vectors": {
                "splade-model-name": {
                    "index": {
                        "on_disk": true
                    }
                }
            }
        }'
