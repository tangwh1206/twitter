#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)

RUNTIME_ROOT=${CURDIR}

export APP_ENV="prod"
export CONF_DIR="$CURDIR/conf"
export LOG_DIR="$RUNTIME_ROOT/log"

if [ ! -d "$LOG_DIR/app" ]; then
    mkdir -p "$LOG_DIR/app"
fi

if [ ! -d "$LOG_DIR/rpc" ]; then
    mkdir -p "$LOG_DIR/rpc"
fi

exec "$CURDIR/bin/twitter"