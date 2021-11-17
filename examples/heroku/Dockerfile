FROM ghcr.io/cased/shell:latest
ADD entrypoint.sh /
ENTRYPOINT [ "/sbin/tini", "--", "/entrypoint.sh" ]
CMD []
RUN apt-get update && \
    apt-get install -y openssh-server vim-nox && \
    adduser --system --disabled-password heroku && \
    curl https://cli-assets.heroku.com/install.sh | sh && \
    echo "PasswordAuthentication No" > /etc/sshd_config

ADD --chown=heroku:nogroup heroku-bashrc /home/heroku/.bashrc
ADD prompts /prompts.json
ENV CASED_SHELL_HOST_FILE=/prompts.json
