from qdrant_client import QdrantClient

client = QdrantClient(url="http://localhost:6333")

client.clear_payload(
    collection_name="{collection_name}",
    points_selector=[0, 3, 100],
)
