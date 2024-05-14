#!/usr/bin/env bash


set -e

# This script will download the latest OpenAPI specification from the
# Qdrant repository and generate the corresponding fern versions of the documents.



PROJECT_ROOT="$(pwd)/$(dirname "$0")/../"

rm -rf qdrant

git clone --sparse --filter=blob:none --depth=1 -b master git@github.com:qdrant/qdrant.git

cd qdrant
git sparse-checkout add docs/redoc

cd $PROJECT_ROOT

# Update master API

cp qdrant/docs/redoc/master/openapi.json $PROJECT_ROOT/fern/apis/master/openapi.json


# Find latest version inside the repository `docs/redoc` starting with `v*`

latest_version=$(ls qdrant/docs/redoc | grep -oP 'v\d.*' | sort -V | tail -n 1)

echo "Latest version found: $latest_version"

# Update latest API specs in `fern/apis`

rm -rf $PROJECT_ROOT/fern/apis/$latest_version

cp -r $PROJECT_ROOT/fern/apis/master $PROJECT_ROOT/fern/apis/$latest_version

cp qdrant/docs/redoc/$latest_version/openapi.json $PROJECT_ROOT/fern/apis/$latest_version/openapi.json



# Create version file in `fern/versions` by replacing master with the latest version

cat $PROJECT_ROOT/fern/versions/master.yml | sed "s/master/$latest_version/g" > $PROJECT_ROOT/fern/versions/$latest_version.yml


# Generate list of all versions to put into docs.yml

#
# versions: 
#  - display-name: master
#    path: ./versions/master.yml
#  - display-name: v1.9.x
#    path: ./versions/v1.9.x.yml 
#
# e.t.c.

VERSIONS_TMP_YAML=qdrant/versions.tmp.yml

echo "versions:" > $VERSIONS_TMP_YAML

echo "  - display-name: master" >> $VERSIONS_TMP_YAML
echo "    path: ./versions/master.yml" >> $VERSIONS_TMP_YAML

for version in $(ls $PROJECT_ROOT/fern/versions | grep -oP 'v\d.*' | sort -V); do
    # cut the .yml extension
    version=$(echo $version | sed 's/\.yml$//g')
    echo "  - display-name: $version" >> $VERSIONS_TMP_YAML
    echo "    path: ./versions/$version.yml" >> $VERSIONS_TMP_YAML
done


# Join the versions with the existing docs.yml using `yq`

./yq eval-all 'select(fileIndex == 0) * select(fileIndex == 1)' $PROJECT_ROOT/fern/docs.yml $VERSIONS_TMP_YAML > $PROJECT_ROOT/fern/docs.tmp.yml

mv $PROJECT_ROOT/fern/docs.tmp.yml $PROJECT_ROOT/fern/docs.yml


