from qdrant_client import QdrantClient, models

client = QdrantClient(url="http://localhost:6333")

client.discover(
    "{collection_name}",
    target=[0.2, 0.1, 0.9, 0.7],
    context=[
        models.ContextExamplePair(positive=100, negative=718),
        models.ContextExamplePair(positive=200, negative=300),
    ],
    limit=10,
)
