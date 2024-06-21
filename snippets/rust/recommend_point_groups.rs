use qdrant_client::qdrant::{RecommendPointGroupsBuilder, RecommendStrategy};
use qdrant_client::Qdrant;

let client = Qdrant::from_url("http://localhost:6334").build()?;

client
    .recommend_groups(
        RecommendPointGroupsBuilder::new("{collection_name}", "document_id", 2, 3)
            .positive(100)
            .positive(200)
            .negative(718)
            .strategy(RecommendStrategy::AverageVector),
    )
    .await?;
