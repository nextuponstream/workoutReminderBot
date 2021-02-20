#!/bin/bash
TIMEOUT=25
BUILD=docker/go
MDB=wrb_mdb:27017
NGDB=wrb_ngdb:7687

wait_for_service()

echo 'Removing all unnecessary containers/images that are used in build'
# yes command: https://stackoverflow.com/a/7642711
yes | docker container prune # for wrb_wait_dbs
docker system prune --volume

./startDbs.sh # dbs started in background

cd ..

rm -rf $BUILD/pkg
cp -R pkg $BUILD/pkg
rm -rf $BUILD/cmd
cp -R cmd $BUILD/cmd
rm $BUILD/go.mod
cp go.mod $BUILD/go.mod

for service in $MDB $NGDB
do
    # wait on dbs to start before running our bots
    docker build --tag wrb_wait_dbs docker/bash
    docker run --network=workoutreminderbot_web -i wrb_wait_dbs bash -c \
        "./wait-for-it.sh -t $TIMEOUT $service"

    hasTimeout=$?

    if [ "$hasTimeout" -ne 0 ]; then # when timeout occurs
        echo "error: mongo database didn't start in time!"
        exit 1
    else
        echo "$service is up"
    fi
done

docker-compose up --build bot
