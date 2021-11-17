FROM ghcr.io/cased/shell:latest
ADD entrypoint.sh /
ENTRYPOINT [ "/sbin/tini", "--", "/entrypoint.sh" ]
CMD []
RUN apt-get update && \
    apt-get install -y openssh-server vim-nox && \
    adduser --disabled-password heroku && \
    curl https://cli-assets.heroku.com/install.sh | sh && \
    echo "PasswordAuthentication No" > /etc/sshd_config && \
    sh - heroku -c "ssh-keygen -P '' -f /home/heroku/.ssh/id_rsa; \
                    cp /home/heroku/.ssh/id_rsa.pub /home/heroku/.ssh/authorized_keys"

ADD --chown=heroku heroku-bashrc /home/heroku/.bashrc
ADD prompts /prompts.json
ENV CASED_SHELL_HOST_FILE=/prompts.json
