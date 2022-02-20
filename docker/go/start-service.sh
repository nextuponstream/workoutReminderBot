#!/bin/bash

echo $SHELL

for SERVICE in $MDB $NGDB
do
    echo $SERVICE
    sh wait-for-it.sh -t $TIMEOUT $SERVICE
    hasTimeout=$?

    if [ "$hasTimeout" -ne 0 ]; then # when timeout occurs
        echo "error: required service $SERVICE didn't start in time!"
        exit 1
    else
        echo "$SERVICE is up"
    fi
done

/bin/workoutReminderBot
