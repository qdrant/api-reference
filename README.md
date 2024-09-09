# Qdrant API Documentation

The website is available at [api.qdrant.tech](https://api.qdrant.tech/).

## What does this repo contain?

- [OpenAPI Specs](./fern/apis/)
- [Docs Configuration](./fern/docs.yml)
- [GitHub Workflows](./.github/workflows)

## How to sync OpenAPI specs?

Either manually, by running the script:

```bash
bash -x tools/sync-openapi.sh
```

> [!NOTE]  
> This script requires GNU `grep`, `yq` and `python` to be present in the system.

Or automatically by creating a PR. The Github Action will automatically try to sync and re-generate all the docs it can.

## How to update snippets?

The source of truth for the snippets is located in `snippets` folder, edit it there.
File names should match the `operationId` of the OpenAPI spec.

After editing, follow the `How to sync OpenAPI specs?` steps to automatically reflect the changes for the latest version of the API. 

## How to deploy the documentation?

To update your documentation, run 

```sh
# npm install -g fern-api 
fern generate --docs
```

To preview your documentation, run 
```sh
# npm install -g fern-api
fern generate --docs --preview
```

The repository contains GitHub workflows that will automatically run 
these commands for you. For example, when you make a PR a preview link 
will be auto-generated and when you merge to main the docs site
will update. 
