async fn main() -> Result<(), qdrant_client::QdrantError> {
    use qdrant_client::qdrant::{Query, QueryBatchPointsBuilder, QueryPointsBuilder};
    use qdrant_client::Qdrant;

    let client = Qdrant::from_url("http://localhost:6334").build()?;

    let request = QueryBatchPointsBuilder::new(
        "{collection_name}",
        vec![
            QueryPointsBuilder::new("")
                .query(Query::new_nearest(vec![0.1; 8]))
                .into(),
            QueryPointsBuilder::new("")
                .query(Query::new_nearest(vec![0.3; 8]))
                .into(),
        ],
    );

    let _response = client.query_batch(request).await?;

    Ok(())
}
