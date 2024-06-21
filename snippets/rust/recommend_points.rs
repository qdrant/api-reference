use qdrant_client::qdrant::{Condition, Filter, RecommendPointsBuilder, RecommendStrategy};
use qdrant_client::Qdrant;

let client = Qdrant::from_url("http://localhost:6334").build()?;

client
    .recommend(
        RecommendPointsBuilder::new("{collection_name}", 3)
            .positive(100)
            .positive(200)
            .positive(vec![100.0, 231.0])
            .negative(718)
            .negative(vec![0.2, 0.3, 0.4, 0.5])
            .strategy(RecommendStrategy::AverageVector)
            .filter(Filter::must([Condition::matches(
                "city",
                "London".to_string(),
            )])),
    )
    .await?;
