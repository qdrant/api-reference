# Qdrant API Documentation

This repo contains the configuration files for the Qdrant's API documentation. 
The documentation is built using Fern. 

Click [here](https://qdrant.docs.buildwithfern.com/api-reference) to see the generated website. 

## What does this repo contain?

- [OpenAPI Specs](./fern/apis/)
- [Docs Configuration](./fern/docs.yml)
- [GitHub Workflows](./.github/workflows)

## How to sync OpenAPI specs?

Either manually, by rinning the script:

```bash
bash -x tools/sync-openapi.sh
```

WARN: This script requires `yq` and `python` to be present in the system.

Or automatically, just create a PR, and preview action will automatically try to sync and re-generate all the docs it can.

## How to update snippets?

The source of truth for the snippets is located in `snippets` folder, edit it there.
File names should match the `operationId` of the OpenAPI spec.

After edited, follow the `How to sync OpenAPI specs?` steps and it will be automatically updated for the latest version of the API. 

## How to deploy documentation live?

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
