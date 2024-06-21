use qdrant_client::Qdrant;

let client = Qdrant::from_url("http://localhost:6334").build()?;

client
    .delete_field_index("{collection_name}", "{field_name}", None)
    .await?;
