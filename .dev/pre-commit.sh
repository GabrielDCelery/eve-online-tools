#!/usr/bin/env bash

docker build -f $(pwd)/docker/Dockerfile.precommit -t eve-online-tools-cli/precommit .
docker run --rm -t -e "TERM=xterm-256color" -v $(pwd):/home/local-user/eve-online-tools-cli eve-online-tools-cli/pre-commit
