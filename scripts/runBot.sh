#!/bin/bash

cd ..
rm -rf docker/go/src
cp -R src docker/go/src # allow for docker to build the app with the src
docker-compose up --build
# FIXME cannot find package my own package