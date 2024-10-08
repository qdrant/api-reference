use qdrant_client::Qdrant;
use qdrant_client::qdrant::{Condition, Filter, PointId, PrefetchQueryBuilder, Query, QueryPointsBuilder};

let client = Qdrant::from_url("http://localhost:6334").build()?;

// Query nearest by ID
let _nearest = client.query(
    QueryPointsBuilder::new("{collection_name}")
        .query(Query::new_nearest(PointId::new("43cf51e2-8777-4f52-bc74-c2cbde0c8b04")))
).await?;

// Recommend on the average of these vectors
let _recommendations = client.query(
    QueryPointsBuilder::new("{collection_name}")
        .query(Query::new_recommend(
            RecommendInputBuilder::default()
                .add_positive(vec![0.1; 8])
                .add_negative(PointId::from(0))
        ))
).await?;

// Fusion query
let _hybrid = client.query(
    QueryPointsBuilder::new("{collection_name}")
        .add_prefetch(PrefetchQueryBuilder::default()
            .query(Query::new_nearest([(1, 0.22), (42, 0.8)].as_slice()))
            .using("sparse")
            .limit(20u64)
        )
        .add_prefetch(PrefetchQueryBuilder::default()
            .query(Query::new_nearest(vec![0.01, 0.45, 0.67]))
            .using("dense")
            .limit(20u64)
        )
        .query(Query::new_fusion(Fusion::Rrf))
).await?;

// 2-stage query
let _refined = client.query(
    QueryPointsBuilder::new("{collection_name}")
        .add_prefetch(PrefetchQueryBuilder::default()
            .query(Query::new_nearest(vec![0.01, 0.45, 0.67]))
            .limit(100u64)
        )
        .query(Query::new_nearest(vec![
            vec![0.1, 0.2],
            vec![0.2, 0.1],
            vec![0.8, 0.9],
        ]))
        .using("colbert")
        .limit(10u64)
).await?;

// Random sampling (as of 1.11.0)
let _sampled = client
    .query(
        QueryPointsBuilder::new("{collection_name}")
            .query(Query::new_sample(Sample::Random))
    )
    .await?;
