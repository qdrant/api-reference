from qdrant_client import QdrantClient

client = QdrantClient(url="http://localhost:6333")

client.delete(
    collection_name="{collection_name}",
    points_selector=models.PointIdsList(
        points=[0, 3, 100],
    ),
)

client.delete(
    collection_name="{collection_name}",
    points_selector=models.FilterSelector(
        filter=models.Filter(
            must=[
                models.FieldCondition(
                    key="color",
                    match=models.MatchValue(value="red"),
                ),
            ],
        )
    ),
)
