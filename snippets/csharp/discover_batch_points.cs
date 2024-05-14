using Qdrant.Client;
using Qdrant.Client.Grpc;

var client = new QdrantClient("localhost", 6334);

var discoverPoints = new List<DiscoverPoints>
{
    new DiscoverPoints
    {
        CollectionName = "{collection_name}",
        Target = new TargetVector
        {
            Single = new VectorExample { Vector = new float[] { 0.2f, 0.1f, 0.9f, 0.7f }, }
        },
        Context =
        {
            new ContextExamplePair()
            {
                Positive = new VectorExample { Id = 100 },
                Negative = new VectorExample { Id = 718 }
            },
            new ContextExamplePair()
            {
                Positive = new VectorExample { Id = 200 },
                Negative = new VectorExample { Id = 300 }
            }
        },
        Limit = 10
    },
    new DiscoverPoints
    {
        CollectionName = "{collection_name}",
        Target = new TargetVector
        {
            Single = new VectorExample { Vector = new float[] { 0.5f, 0.3f, 0.2f, 0.3f }, }
        },
        Context =
        {
            new ContextExamplePair()
            {
                Positive = new VectorExample { Id = 342 },
                Negative = new VectorExample { Id = 213 }
            },
            new ContextExamplePair()
            {
                Positive = new VectorExample { Id = 100 },
                Negative = new VectorExample { Id = 200 }
            }
        },
        Limit = 10
    }
};
await client.DiscoverBatchAsync("{collection_name}", discoverPoints);
