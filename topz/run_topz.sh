#!/bin/sh
APP_ID='c80d17f5f638'
docker run --pid=container:${APP_ID} \
        -p 8083:8083 brendanburns/topz:db0fa58 /server -addr 0.0.0.0:8083

