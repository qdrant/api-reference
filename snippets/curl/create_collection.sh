curl -X PUT http://localhost:6333/collections/collection_name \
     -H "api-key: <apiKey>" \
     -H "Content-Type: application/json" \
     -d '{
            "vectors": {
                "size": 300,
                "distance": "Cosine"
            } 
        }'
