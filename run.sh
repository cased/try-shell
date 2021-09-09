#!/bin/bash

set -e

tmpdir=$(mktemp -d)
trap "rm -rf $tmpdir" EXIT

if grep -q try-shell docker-compose.yml>/dev/null; then
  cp docker-compose.yml $tmpdir/docker-compose.yml
else
  : # TODO obtain docker-compose.yml via curl once project is OSS'd
fi

echo "Starting Cased Shell. Press Ctrl+C to quit" 1>&2

docker compose --project-name try-shell --file $tmpdir/docker-compose.yml up \
  --always-recreate-deps --force-recreate --remove-orphans --renew-anon-volumes
