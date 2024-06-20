from qdrant_client import QdrantClient, models

client = QdrantClient(url="http://localhost:6333")

client.search(
    collection_name="{collection_name}",
    query_filter=models.Filter(
        must=[
            models.FieldCondition(
                key="city",
                match=models.MatchValue(
                    value="London",
                ),
            )
        ]
    ),
    search_params=models.SearchParams(hnsw_ef=128, exact=False),
    query_vector=[0.2, 0.1, 0.9, 0.7],
    limit=3,
)