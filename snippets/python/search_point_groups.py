from qdrant_client import QdrantClient

client = QdrantClient(url="http://localhost:6333")

client.search_groups(
    collection_name="{collection_name}",
    query_vector=[1.1],
    group_by="document_id",
    limit=4,
    group_size=2,
)
