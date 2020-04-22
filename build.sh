#!/bin/bash
docker run --rm -it -v $PWD:/app -w /app golang go build ipfs-block-to-cid.go 
