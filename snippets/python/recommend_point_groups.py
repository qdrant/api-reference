from qdrant_client import QdrantClient, models

client = QdrantClient(url="http://localhost:6333")

client.recommend_groups(
    collection_name="{collection_name}",
    positive=[100, 231],
    negative=[718],
    group_by="document_id",
    limit=3,
    group_size=2,
)
