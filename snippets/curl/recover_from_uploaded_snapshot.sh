curl  -X POST \
  'http://localhost:6333/collections/collection_name/snapshots/upload' \
  --header 'api-key: <api-key-value>' \
  --form 'snapshot=@/path/to/snapshot.shapshot'