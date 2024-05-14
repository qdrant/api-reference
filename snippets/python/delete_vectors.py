from qdrant_client import QdrantClient

client = QdrantClient(url="http://localhost:6333")

client.delete_vectors(
    collection_name="{collection_name}",
    points=[0, 3, 100],
    vectors=["text", "image"],
)
