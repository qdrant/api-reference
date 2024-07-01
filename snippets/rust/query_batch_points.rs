use qdrant_client::{Qdrant, QdrantError};
use qdrant_client::qdrant::{Query, QueryPointsBuilder, QueryBatchPointsBuilder};

let request = QueryBatchPointsBuilder::new("{collection_name}", vec![
    QueryPointsBuilder::new("")
        .query(Query::new_nearest(
            vec![0.1; 8],
        )).into(),
    QueryPointsBuilder::new("")
        .query(Query::new_nearest(
            vec![0.3; 8],
        )).into(),
]);

let response = client.query_batch(request).await?;