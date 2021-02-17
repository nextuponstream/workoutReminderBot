#!/bin/bash
TIMEOUT=25

./startDbs.sh # dbs started in background

BUILD=docker/go

// TODO remove
echo 'Removing all unnecessary containers/images that are used in build'
docker system prune --volume

cd ..

rm -rf $BUILD/pkg
cp -R pkg $BUILD/pkg
rm -rf $BUILD/cmd
cp -R cmd $BUILD/cmd
rm $BUILD/go.mod
cp go.mod $BUILD/go.mod

# wait on dbs to start before running our bots
docker build --tag wrb_wait_dbs docker/bash
docker run --network=workoutreminderbot_web -i wrb_wait_dbs bash -c \
    "./wait-for-it.sh -t $TIMEOUT wrb_mdb:27017"

hasTimeout=$?

if [ "$hasTimeout" -ne 0 ]; then # when timeout occurs
    echo "error: mongo database didn't start in time!"
    exit 1
fi

// TODO wait on neo4j gdb

docker-compose up --build bot
