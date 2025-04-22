import { QdrantClient } from "@qdrant/js-client-rest";

const client = new QdrantClient({ host: "localhost", port: 6333 });

// Query nearest by ID
let _nearest = client.query("{collection_name", {
    query: "43cf51e2-8777-4f52-bc74-c2cbde0c8b04"
});

// Recommend on the average of these vectors
let _recommendations = client.query("{collection_name}", {
    query: {
        recommend: {
            positive: ["43cf51e2-8777-4f52-bc74-c2cbde0c8b04", [0.11, 0.35, 0.6]],
            negative: [0.01, 0.45, 0.67]
        }
    }
});

// Fusion query
let _hybrid = client.query("{collection_name}", {
    prefetch: [
        {
            query: {
                values: [0.22, 0.8],
                indices: [1, 42],
            },
            using: 'sparse',
            limit: 20,
        },
        {
            query: [0.01, 0.45, 0.67],
            using: 'dense',
            limit: 20,
        },
    ],
    query: {
        fusion: 'rrf',
    },
});

// 2-stage query
let _refined = client.query("{collection_name}", {
    prefetch: {
        query: [1, 23, 45, 67],
        limit: 100,
    },
    query: [
        [0.1, 0.2],
        [0.2, 0.1],
        [0.8, 0.9],
    ],
    using: 'colbert',
    limit: 10,
});

// Random sampling (as of 1.11.0)
let _sampled = client.query("{collection_name}", {
  query: { sample: "random" },
});

// Score boost depending on payload conditions (as of 1.14.0)
const tag_boosted = await client.query("{collection_name}", {
  prefetch: {
    query: [0.2, 0.8, 0.1, 0.9],
    limit: 50
  },
  query: {
    formula: {
      sum: [
        "$score",
        {
          mult: [ 0.5, { key: "tag", match: { any: ["h1", "h2", "h3", "h4"] }} ]
        },
        {
          mult: [ 0.25, { key: "tag", match: { any: ["p", "li"] }} ]
        }
      ]
    }
  }
});

// Score boost geographically closer points (as of 1.14.0)
const distance_boosted = await client.query("{collection_name}", {
  prefetch: {
    query: [0.2, 0.8, ...],
    limit: 50
  },
  query: {
    formula: {
      sum: [
        "$score",
        {
          gauss_decay: {
            x: {
              geo_distance: {
                origin: { lat: 52.504043, lon: 13.393236 }, // Berlin
                to: "geo.location"
              }
            },
            scale: 5000 // 5km
          }
        }
      ]
    },
    defaults: { "geo.location": { lat: 48.137154, lon: 11.576124 } } // Munich
  }
});
