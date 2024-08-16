from qdrant_client import QdrantClient, models

client = QdrantClient(url="http://localhost:6333")

# Query nearest by ID
nearest = client.query_points(
    collection_name="{collection_name}",
    query="43cf51e2-8777-4f52-bc74-c2cbde0c8b04",
)

# Recommend on the average of these vectors
recommended = client.query_points(
    collection_name="{collection_name}",
    query=models.RecommendQuery(recommend=models.RecommendInput(
        positive=["43cf51e2-8777-4f52-bc74-c2cbde0c8b04", [0.11, 0.35, 0.6, ...]],
        negative=[[0.01, 0.45, 0.67, ...]]
    ))
)

# Fusion query
hybrid = client.query_points(
    collection_name="{collection_name}",
    prefetch=[
        models.Prefetch(
            query=models.SparseVector(indices=[1, 42], values=[0.22, 0.8]),
            using="sparse",
            limit=20,
        ),
        models.Prefetch(
            query=[0.01, 0.45, 0.67, ...],  # <-- dense vector
            using="dense",
            limit=20,
        ),
    ],
    query=models.FusionQuery(fusion=models.Fusion.RRF),
)

# 2-stage query
refined = client.query_points(
    collection_name="{collection_name}",
    prefetch=models.Prefetch(
        query=[0.01, 0.45, 0.67, ...],  # <-- dense vector
        limit=100,
    ),
    query=[
        [0.1, 0.2, ...],  # <─┐
        [0.2, 0.1, ...],  # < ├─ multi-vector
        [0.8, 0.9, ...],  # < ┘
    ],
    using="colbert",
    limit=10,
)

# Random sampling (as of 1.11.0)
sampled = client.query_points(
    collection_name="{collection_name}",
    query=models.SampleQuery(sample=models.Sample.Random)
)
