from qdrant_client import QdrantClient

client = QdrantClient(url="http://localhost:6333")

client.recover_snapshot(
    "{collection_name}",
    "http://example.com/path/to/snapshot.shapshot",
)
