#!/usr/bin/env bash

docker-compose -f "meta.docker-compose.yaml" up -d --build
docker-compose -f "meta.docker-compose.yaml" run test_generate
docker-compose -f "meta.docker-compose.yaml" run generate
