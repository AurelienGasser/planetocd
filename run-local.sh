#!/bin/bash

set -eu

trap "trap - SIGTERM && kill -- -$$" SIGINT SIGTERM EXIT

which nodemon >/dev/null 2>/dev/null || (echo you need to install nodemon with "npm install -g nodemon" && exit 1)
which CompileDaemon >/dev/null 2>/dev/null || (echo you need to install CompileDeamon with "go [get|install] github.com/githubnemo/CompileDaemon" && exit 1)
which psql >/dev/null 2>/dev/null || (echo you need to install psql && exit 1)

psql "${PLANETOCD_DATABASE_URL}" -c "\
CREATE TABLE IF NOT EXISTS likes(\
    id SERIAL PRIMARY KEY, \
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, \
    article_id INT NOT NULL, \
    ip VARCHAR(255), \
    username VARCHAR(255), \
    random_number INT NOT NULL\
);"
nodemon -w "./server/static/css" -e "less" --exec "make css" &
CompileDaemon -directory=. -build="go build -o ./bin/planetocd" -command="./bin/planetocd"  -pattern='(.+\.go|.+\.c|.+\.html|.+\.yaml|.+\.css|.+\.json|\.md)$' &

wait
