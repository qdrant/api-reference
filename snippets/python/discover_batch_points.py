from qdrant_client import QdrantClient, models

client = QdrantClient(url="http://localhost:6333")

discover_queries = [
    models.DiscoverRequest(
        target=[0.2, 0.1, 0.9, 0.7],
        context=[
            models.ContextExamplePair(
                positive=100,
                negative=718,
            ),
            models.ContextExamplePair(
                positive=200,
                negative=300,
            ),
        ],
        limit=10,
    ),
    models.DiscoverRequest(
        target=[0.5, 0.3, 0.2, 0.3],
        context=[
            models.ContextExamplePair(
                positive=342,
                negative=213,
            ),
            models.ContextExamplePair(
                positive=100,
                negative=200,
            ),
        ],
        limit=5,
    ),
]

client.discover_batch("{collection_name}", discover_queries)
