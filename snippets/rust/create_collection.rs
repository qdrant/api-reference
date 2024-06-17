use qdrant_client::{
    client::QdrantClient,
    qdrant::{vectors_config::Config, CreateCollection, Distance, VectorParams, VectorsConfig},
};

let client = QdrantClient::from_url("http://localhost:6334").build()?;

client
    .create_collection(&CreateCollection {
        collection_name: "{collection_name}".to_string(),
        vectors_config: Some(VectorsConfig {
            config: Some(Config::Params(VectorParams {
                size: 100,
                distance: Distance::Cosine.into(),
                ..Default::default()
            })),
        }),
        ..Default::default()
    })
    .await?;

//  Or with a sparse vector field

client
    .create_collection(&CreateCollection {
        collection_name: "{collection_name}".to_string(),
        sparse_vectors_config: Some(SparseVectorConfig {
            map: [(
                "text".to_string(),
                SparseVectorParams {
                    ..Default::default()
                },
            )]
            .into(),
            ..Default::default()
        }),
        ..Default::default()
    })
    .await?;