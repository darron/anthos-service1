#!/bin/sh

while [ 1 ]; do
    curl http://service1.default.svc.cluster.local:3000/
    sleep 1
done