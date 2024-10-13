from qdrant_client import QdrantClient, models

client = QdrantClient(url="http://localhost:6333")

client.facet(
    collection_name="{collection_name}",
    key="my-payload-key",
    facet_filter=models.Filter(must=[models.Match("color", "red")]),
    limit=10,
)

