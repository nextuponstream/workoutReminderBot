#!/bin/bash
TIMEOUT=25

./startDbs.sh # dbs started in background

cd ..
rm -rf docker/go/src
cp -R src docker/go/src # allow for docker to build the app with the src

# wait on dbs to start before running our bots
docker build --tag wrb_wait_dbs docker/bash
docker run --network=workoutreminderbot_web -i wrb_wait_dbs bash -c \
    "./wait-for-it.sh -t $TIMEOUT wrb_mdb:27017"

hasTimeout=$?

if [ "$hasTimeout" -ne 0 ]; then # when timeout occurs
    echo "error: mongo database didn't start in time!"
    exit 1
fi

docker-compose up --build bot
