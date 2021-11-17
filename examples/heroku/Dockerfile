FROM ghcr.io/cased/shell:latest
ADD entrypoint.sh /
ENTRYPOINT [ "/bin/bash", "/entrypoint.sh" ]
CMD []
RUN apt-get update && \
    apt-get install -y openssh-server vim-nox && \
    curl https://cli-assets.heroku.com/install.sh | sh && \
    echo "PasswordAuthentication No" > /etc/sshd_config

ADD /heroku-bashrc
ADD prompts.json /
ENV CASED_SHELL_HOST_FILE=/prompts.json
