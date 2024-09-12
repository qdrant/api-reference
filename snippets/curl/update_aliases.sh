# Create an alias
curl  -X POST \
  'http://localhost:6333/collections/aliases' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "actions": [
    {
      "create_alias": {
        "collection_name": "{collection_name}",
        "alias_name": "{alias_name}"
      }
    }
  ]
}'

# Delete an alias
curl  -X POST \
  'http://localhost:6333/collections/aliases' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "actions": [
    {
      "delete_alias": {
        "alias_name": "{alias_name}"
      }
    }
  ]
}'

# Rename an alias
curl  -X POST \
  'http://localhost:6333/collections/aliases' \
  --header 'api-key: <api-key-value>' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "actions": [
    {
      "rename_alias": {
        "old_alias_name": "{old_alias_name}",
        "new_alias_name": "{new_alias_name}"
      }
    }
  ]
}'
