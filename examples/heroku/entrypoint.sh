#!/bin/bash

# Configure for Heroku
export CASED_SHELL_PORT=$PORT
export CASED_SHELL_TLS=off
: ${CASED_SHELL_LOG_LEVEL:="error"}

python -u run.py --logging=$CASED_SHELL_LOG_LEVEL