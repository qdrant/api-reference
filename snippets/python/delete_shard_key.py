from qdrant_client import QdrantClient

client = QdrantClient(url="http://localhost:6333")

client.delete_shard_key("{collection_name}", "{shard_key}")
