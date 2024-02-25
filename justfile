_default:
  @just --list --unsorted

set dotenv-load

HERE           := justfile_directory()
COMPOSE_FILE   := HERE / "docker-compose.yml"
DOCKER_COMPOSE := "docker compose --file " + COMPOSE_FILE

build:
  echo "Building Docker image"
  docker build --tag $MAINTAINER/$PROJECT:$VERSION {{HERE}}

start: build
 echo "TODO"

delete: #stop
  -docker rm --force $LOGSTEALER_CONTAINER

purge: && delete
  -docker image rm --force $MAINTAINER/$PROJECT:$VERSION
