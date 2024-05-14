from qdrant_client import QdrantClient

client = QdrantClient(url="http://localhost:6333")

client.update_collection(
    collection_name="{collection_name}",
    optimizer_config=models.OptimizersConfigDiff(indexing_threshold=10000),
)
