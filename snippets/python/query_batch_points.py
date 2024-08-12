from qdrant_client import QdrantClient, models

client = QdrantClient(url="http://localhost:6333")

nearest = client.query_batch_points(
    collection_name="{collection_name}",
    requests=[
        models.QueryRequest(
            query=[0.01, 0.45, 0.67],
        ),
        models.QueryRequest(
            query=[0.11, 0.35, 0.6],
        ),
    ]
)