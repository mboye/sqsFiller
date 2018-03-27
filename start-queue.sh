#!/bin/bash
set -x
docker run -p 9324:9324 -v "$(pwd):/etc/elasticmq" s12v/elasticmq
