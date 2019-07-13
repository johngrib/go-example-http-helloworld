#! /usr/bin/env bash

if [[ "$1" = "build" ]]; then
    docker build -t hello:latest .

    exit;
elif [[ "$1" = "up" ]]; then
     docker run -d -t -p 3000:3000 hello

    exit;
fi
