from qdrant_client import QdrantClient

client = QdrantClient(url="http://localhost:6333")

client.list_full_snapshots()
