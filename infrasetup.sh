#!/bin/bash

docker_net_setup() {
  local name="$1"
  if ! docker network inspect "$name" >/dev/null 2>&1
  then
    docker network create "$name"
  fi
}


docker_net_setup ad


docker_run() {
  local name="$1"
  local net="$2"
  shift 2

  if ! docker inspect --type container "$name" >/dev/null 2>&1
  then
    docker run \
        -d \
        --name "$name" \
        --net "$net" \
        "$@"
  fi
}


#docker_run c1 ad consul agent --dev -client 0.0.0.0
docker_run c1 ad consul agent -server -bootstrap -client=0.0.0.0
docker_run c2 ad consul agent -server -join=c1 -client=0.0.0.0
docker_run c3 ad consul agent -server -join=c1 -client=0.0.0.0

docker_run my ad -e MYSQL_ALLOW_EMPTY_PASSWORD=true mysql
