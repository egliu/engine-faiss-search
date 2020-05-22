#!/usr/bin/env bash

set -euxo pipefail

docker rm -f efs kafka zk || true

echo "start local dev env"

docker run -d --name=zk --rm \
  -p 127.0.0.1:2181:2181 \
  wurstmeister/zookeeper

docker run -d --name=kafka --rm \
  --publish 9092:9092 \
  --link zk \
  --env KAFKA_ZOOKEEPER_CONNECT=127.0.0.1:2181 \
  --env KAFKA_ADVERTISED_HOST_NAME=127.0.0.1 \
  --env KAFKA_ADVERTISED_PORT=9092  \
  wurstmeister/kafka

code_root=/go/src/engine-faiss-search

docker run -d --name=efs \
  --network=host \
  --add-host="zk:127.0.0.1" \
  --add-host="kafka:127.0.0.1" \
  -e TZ=Asia/Shanghai \
  -v ${PWD}:${code_root} \
  -w ${code_root} \
  ${baseImage} bash -c "while true; do sleep 1000; done"

docker exec \
  -e COLUMNS="$(tput cols)" \
  -e LINES="$(tput lines)" \
  -it efs bash