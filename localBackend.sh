#!/bin/sh

cd backend && go build -o local cmd/server/main.go 
MYSQL_SERVER=127.0.0.1 \
    MYSQL_PORT=3306 
    MYSQL_USER=itinerary \
    MYSQL_PASSWORD=password \
    MYSQL_DATABASE=itinerary \
    ./local
