#!/bin/sh

cd frontend && rm -rf dist/ && npm run build
docker compose build --no-cache
docker compose up --force-recreate -d