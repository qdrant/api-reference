from qdrant_client import QdrantClient

client = QdrantClient(url="http://localhost:6333")

client.collection_exists(collection_name="{collection_name}")