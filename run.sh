#!/bin/bash

set -e

tmpdir=$(mktemp -d)
trap "docker compose --project-name try-shell --file $tmpdir/docker-compose.yml down; rm -rf $tmpdir" EXIT

if grep -q try-shell docker-compose.yml>/dev/null; then
  cp docker-compose.yml $tmpdir/docker-compose.yml
else
  curl https://github.com/cased/try-shell/raw/main/docker-compose.yml > $tmpdir/docker-compose.yml
fi

if [ -z "${CASED_SHELL_HOSTNAME}" ]; then
  echo "CASED_SHELL_HOSTNAME required" 1>&2
  exit 1
fi
export CASED_SHELL_PORT=$(echo "${CASED_SHELL_HOSTNAME}" | cut -f 2 -d :)

echo "Starting Cased Shell. Press Ctrl+C to quit and cleanup" 1>&2

docker compose --project-name try-shell --file $tmpdir/docker-compose.yml up \
  --always-recreate-deps --force-recreate --remove-orphans --renew-anon-volumes
