from qdrant_client import QdrantClient, models

client = QdrantClient(url="http://localhost:6333")

client.create_collection(
    collection_name="{collection_name}",
    vectors_config=models.VectorParams(size=100, distance=models.Distance.COSINE),
)

# Or with a sparse vector field

client.create_collection(
    collection_name="{collection_name}",
    sparse_vectors_config={
        "text": models.SparseVectorParams(),
    },
)