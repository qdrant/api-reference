from qdrant_client import QdrantClient, models

client = QdrantClient(url="http://localhost:6333")

client.recommend(
    collection_name="{collection_name}",
    positive=[100, 231],
    negative=[718, [0.2, 0.3, 0.4, 0.5]],
    strategy=models.RecommendStrategy.AVERAGE_VECTOR,
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
    limit=3,
)
