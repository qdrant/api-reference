from qdrant_client import QdrantClient

client = QdrantClient(url="http://localhost:6333")

client.delete_payload_index("{collection_name}", "{field_name}");
