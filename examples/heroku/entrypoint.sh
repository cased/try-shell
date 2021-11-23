#!/bin/bash

# Configure for Heroku
export CASED_SHELL_PORT=$PORT
export CASED_SHELL_TLS=off
export CASED_SHELL_HOSTNAME=$HEROKU_APP_NAME.herokuapp.com
: ${CASED_SHELL_LOG_LEVEL:="error"}
export CASED_SHELL_OAUTH_UPSTREAM=localhost:2225

/bin/heroku-ssh http://localhost:$PORT &
python -u run.py --logging=$CASED_SHELL_LOG_LEVEL &

wait -n