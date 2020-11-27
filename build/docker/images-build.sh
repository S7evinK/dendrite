#!/bin/bash

cd $(git rev-parse --show-toplevel)

TAG=${1:-latest}

echo "Building tag '${TAG}'"

docker build -f build/docker/Dockerfile -t matrixdotorg/dendrite:${TAG} .

#docker build -t registry.s3cr3t.me/s7evink/dendrite-monolith:${TAG}  -f build/docker/Dockerfile.monolith .
docker build -t registry2.s3cr3t.me/s7evink/dendrite-polylith:${TAG}  -f build/docker/Dockerfile.polylith .

docker build -t registry2.s3cr3t.me/s7evink/dendrite:generatekeys      --build-arg component=generate-keys                     -f build/docker/Dockerfile.generatekeys .