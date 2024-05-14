use qdrant_client::{client::QdrantClient, qdrant::PointStruct};
use serde_json::json;

let client = QdrantClient::from_url("http://localhost:6334").build()?;

client
    .upsert_points_batch_blocking(
        "{collection_name}".to_string(),
        None,
        vec![
            PointStruct::new(
                1,
                vec![0.9, 0.1, 0.1],
                json!(
                    {"color": "red"}
                )
                .try_into()
                .unwrap(),
            ),
            PointStruct::new(
                2,
                vec![0.1, 0.9, 0.1],
                json!(
                    {"color": "green"}
                )
                .try_into()
                .unwrap(),
            ),
            PointStruct::new(
                3,
                vec![0.1, 0.1, 0.9],
                json!(
                    {"color": "blue"}
                )
                .try_into()
                .unwrap(),
            ),
        ],
        None,
        100,
    )
    .await?;
