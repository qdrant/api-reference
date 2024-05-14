from qdrant_client import QdrantClient, models

client = QdrantClient(url="http://localhost:6333")

filter_ = models.Filter(
    must=[
        models.FieldCondition(
            key="city",
            match=models.MatchValue(
                value="London",
            ),
        )
    ]
)

search_queries = [
    models.SearchRequest(vector=[0.2, 0.1, 0.9, 0.7], filter=filter_, limit=3),
    models.SearchRequest(vector=[0.5, 0.3, 0.2, 0.3], filter=filter_, limit=3),
]

client.search_batch(collection_name="{collection_name}", requests=search_queries)
