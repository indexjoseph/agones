#!/usr/bin/env bash

docker system prune -a --volumes
docker rm -vf $(docker ps -aq)
docker rmi -f $(docker images -aq)
make gen-install
make build-images
make gen-crd-code
make push
make install
