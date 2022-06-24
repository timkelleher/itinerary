#!/bin/sh

docker run \
    --name itinerary_local_db \
    -e MYSQL_ROOT_PASSWORD=root \
    -e MYSQL_DATABASE=itinerary \
    -e MYSQL_USER=itinerary \
    -e MYSQL_PASSWORD=password \
    -p 3306:3306 -d \
    docker.io/library/mariadb:latest