#!/bin/bash
set -xu -p pipefail


export MYSQL_HOST=${MYSQL_HOST:-127.0.0.1}
export MYSQL_PORT=${MYSQL_PORT:-3306}
export MYSQL_USER=${MYSQL_USER:-user}
export LANG="C.UTF-8"
export MYSQL_PWD=${MYSQL_PASS:-password}

cd /docker-entrypoint-initdb.d
mysql -h $MYSQL_HOST -P $MYSQL_PORT -u $MYSQL_USER -p$MYSQL_PWD < 00_create.sql