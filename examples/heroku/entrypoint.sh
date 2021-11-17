#!/bin/bash -ex

adduser --system --disabled-password heroku
su - heroku -c "ssh-keygen -P '' -f /home/heroku/.ssh/id_rsa; \
                cp /home/heroku/.ssh/id_rsa.pub /home/heroku/.ssh/authorized_keys"
mv /heroku-bashrc /home/heroku/.bashrc
chown heroku /home/heroku/.bashrc

/usr/sbin/sshd &

# Configure for Heroku
export CASED_SHELL_PORT=$PORT
export CASED_SHELL_TLS=off
export CASED_SHELL_HOSTNAME=$HEROKU_APP_NAME.herokuapp.com
: ${CASED_SHELL_LOG_LEVEL:="error"}
export CASED_SHELL_SSH_PRIVATE_KEY="$(cat /home/heroku/.ssh/id_rsa)"

python -u run.py --logging=$CASED_SHELL_LOG_LEVEL &

wait -n