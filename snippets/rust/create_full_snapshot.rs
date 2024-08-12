async fn main() -> Result<(), qdrant_client::QdrantError> {
    use qdrant_client::Qdrant;

    let client = Qdrant::from_url("http://localhost:6334").build()?;

    client.create_full_snapshot().await?;
    Ok(())
}
