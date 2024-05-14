from qdrant_client import QdrantClient

client = QdrantClient(url="http://localhost:6333")

client.retrieve(
    collection_name="{collection_name}",
    ids=[0, 3, 100],
)
