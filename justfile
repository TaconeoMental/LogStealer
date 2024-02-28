_default:
  @just --list --unsorted

HERE           := justfile_directory()
COMPOSE_FILE   := HERE / "docker-compose.yml"
DOCKER_COMPOSE := "docker compose --file " + COMPOSE_FILE

set dotenv-load

[private]
create_volume:
    #!/usr/bin/env bash
    set -euo pipefail
    echo -n "[*] Volume '$DB_VOLUME' "
    # Robado de mí mismo jaja:
    # https://github.com/TaconeoMental/sigint.in/blob/f586f8511e87bb82883ae4cea00b57c0e18423c8/sigint.in-docker.sh#L11
    vol_info=$(docker volume ls -f name=$DB_VOLUME --format json | awk '{print $NF}')
    if [ $vol_info ]
    then
        echo "already exists"
    else
        echo "does not exist. Creating..."
        docker volume create $DB_VOLUME
    fi

build:
    echo "Building Docker image"
    docker build --tag $MAINTAINER/$PROJECT:$VERSION {{HERE}}

start: build create_volume
    {{DOCKER_COMPOSE}} up

delete: #stop
    -{{DOCKER_COMPOSE}} down
    -docker network rm --force logstealer_network
    -docker volume rm --force $DB_VOLUME

purge: delete
    -docker image rm --force $MAINTAINER/$PROJECT:$VERSION
