#!/usr/bin/env bash

set -euxo pipefail

docker rm -f efs || true

echo "start local dev env"

code_root=/go/src/engine-faiss-search

docker run -d --name=efs \
  --network=host \
  -e TZ=Asia/Shanghai \
  -v ${PWD}:${code_root} \
  -w ${code_root} \
  ${baseImage} bash -c "while true; do sleep 1000; done"

docker exec \
  -e COLUMNS="$(tput cols)" \
  -e LINES="$(tput lines)" \
  -it efs bash