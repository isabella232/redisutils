#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
ROOT=$( dirname $DIR )

pushd $ROOT > /dev/null
docker build -t cflondonservices/redisutils .
popd > /dev/null
