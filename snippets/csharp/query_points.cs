using Qdrant.Client;
using Qdrant.Client.Grpc;

var client = new QdrantClient("localhost", 6334);

// Query nearest by ID
await client.QueryAsync(
	collectionName: "{collection_name}",
	query: Guid.Parse("43cf51e2-8777-4f52-bc74-c2cbde0c8b04")
);

// Recommend on the average of these vectors
await client.QueryAsync(
	collectionName: "{collection_name}",
	query: new RecommendInput
	{
		Positive =
		{
			Guid.Parse("43cf51e2-8777-4f52-bc74-c2cbde0c8b04"),
			new float[] { 0.11f, 0.35f, 0.6f }
		},
		Negative = { new float[] { 0.01f, 0.45f, 0.67f } }
	}
);

// Fusion query
await client.QueryAsync(
	collectionName: "{collection_name}",
	prefetch: new List<PrefetchQuery>
	{
		new()
		{
			Query = new (float, uint)[] { (0.22f, 1), (0.8f, 42), },
			Using = "sparse",
			Limit = 20
		},
		new()
		{
			Query = new float[] { 0.01f, 0.45f, 0.67f },
			Using = "dense",
			Limit = 20
		}
	},
	query: Fusion.Rrf
);

// 2-stage query
await client.QueryAsync(
	collectionName: "{collection_name}",
	prefetch: new List<PrefetchQuery>
	{
		new() { Query = new float[] { 0.01f, 0.45f, 0.67f }, Limit = 100 }
	},
	query: new float[][] { [0.1f, 0.2f], [0.2f, 0.1f], [0.8f, 0.9f] },
	usingVector: "colbert",
	limit: 10
);

// Random sampling (as of 1.11.0)
await client.QueryAsync(
    collectionName: "{collection_name}",
    query: Sample.Random
);

// Score boost depending on payload conditions (as of 1.14.0)
await client.QueryAsync(
	collectionName: "{collection_name}",
	prefetch:
	[
		new PrefetchQuery { Query = new float[] { 0.01f, 0.45f, 0.67f }, Limit = 100 },
	],
	query: new Formula
	{
		Expression = new SumExpression
		{
			Sum =
			{
				"$score",
				new MultExpression
				{
					Mult = { 0.5f, Match("tag", ["h1", "h2", "h3", "h4"]) },
				},
				new MultExpression { Mult = { 0.25f, Match("tag", ["p", "li"]) } },
			},
		},
	},
	limit: 10
);

// Score boost geographically closer points (as of 1.14.0)
await client.QueryAsync(
	collectionName: "{collection_name}",
	prefetch:
	[
		new PrefetchQuery { Query = new float[] { 0.01f, 0.45f, 0.67f }, Limit = 100 },
	],
	query: new Formula
	{
		Expression = new SumExpression
		{
			Sum =
			{
				"$score",
				WithExpDecay(
					new()
					{
						X = new GeoDistance
						{
							Origin = new GeoPoint { Lat = 52.504043, Lon = 13.393236 },
							To = "geo.location",
						},
						Scale = 5000,
					}
				),
			},
		},
		Defaults =
		{
			["geo.location"] = new Dictionary<string, Value>
			{
				["lat"] = 48.137154,
				["lon"] = 11.576124,
			},
		},
	}
);
