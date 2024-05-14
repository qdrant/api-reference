use qdrant_client::qdrant::{
    points_selector::PointsSelectorOneOf,
    points_update_operation::{
        Operation, PointStructList, SetPayload, UpdateVectors,
    },
    PointStruct, PointVectors, PointsIdsList, PointsSelector, PointsUpdateOperation,
};

let client = QdrantClient::from_url("http://localhost:6334").build()?;

client
    .batch_updates_blocking(
        "{collection_name}",
        &[
            PointsUpdateOperation {
                operation: Some(Operation::Upsert(PointStructList {
                    points: vec![PointStruct::new(
                        1,
                        vec![1.0, 2.0, 3.0, 4.0],
                        json!({}).try_into().unwrap(),
                    )],
                    ..Default::default()
                })),
            },
            PointsUpdateOperation {
                operation: Some(Operation::UpdateVectors(UpdateVectors {
                    points: vec![PointVectors {
                        id: Some(1.into()),
                        vectors: Some(vec![1.0, 2.0, 3.0, 4.0].into()),
                    }],
                    ..Default::default()
                })),
            },
            PointsUpdateOperation {
                operation: Some(Operation::OverwritePayload(SetPayload {
                    points_selector: Some(PointsSelector {
                        points_selector_one_of: Some(PointsSelectorOneOf::Points(
                            PointsIdsList {
                                ids: vec![1.into()],
                            },
                        )),
                    }),
                    payload: HashMap::from([("test_payload".to_string(), 1.into())]),
                    ..Default::default()
                })),
            },
        ],
        None,
    )
    .await?;
