from qdrant_client import QdrantClient

client = QdrantClient(url="http://localhost:6333")

client.delete_payload(
    collection_name="{collection_name}",
    keys=["color", "price"],
    points=[0, 3, 100],
)
