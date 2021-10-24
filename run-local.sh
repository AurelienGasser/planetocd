#!/bin/bash

trap "trap - SIGTERM && kill -- -$$" SIGINT SIGTERM EXIT

which nodemon >/dev/null 2>/dev/null || (echo you need to install nodemon with "npm install -g nodemon" && exit 1)
which CompileDaemon >/dev/null 2>/dev/null || (echo you need to install CompileDeamon with "go get github.com/githubnemo/CompileDaemon" && exit 1)

nodemon -w "./static/css" -e "less" --exec "make css" &
CompileDaemon -directory=. -command="./planetocd"  -pattern='(.+\.go|.+\.c|.+\.html|.+\.yaml|.+\.css|.+\.json|\.md)$' &

wait
