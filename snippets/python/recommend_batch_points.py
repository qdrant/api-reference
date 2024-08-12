from qdrant_client import QdrantClient, models

client = QdrantClient(url="http://localhost:6333")

filter_ = models.Filter(must=[models.FieldCondition(key="city", match=models.MatchValue(value="London"))])

recommend_queries = [
    models.RecommendRequest(
        positive=[100, 231], negative=[718], filter=filter_, limit=3
    ),
    models.RecommendRequest(positive=[200, 67], negative=[300], limit=3),
]

client.recommend_batch(collection_name="{collection_name}", requests=recommend_queries)
