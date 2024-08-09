# Qdrant API Documentation

This repo contains the configuration files for Qdrant's API documentation. 
The documentation is built using Fern. 

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

Or automatically, just create a PR, and preview action will automatically try to sync and re-generate all the docs it can.

## How to update snippets?

The source of truth for the snippets is located in `snippets` folder, edit it there.
File names should match the `operationId` of the OpenAPI spec.

After editing, follow the `How to sync OpenAPI specs?` steps to automatically reflect the changes for the latest version of the API. 


## How to preview changes?

To check the changes locally with hot-reload, run 

```sh
# npm install -g fern-api 
fern docs dev
```

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

## LSP support on snippets

### Rust

The rust snippets project is described at `Cargo.toml`, and each snippet is a unique module at `snippets/rust/lib.rs`

### Python

Just make sure to `$ pip install qdrant_client` in the current environment.
