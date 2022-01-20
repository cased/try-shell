FROM ghcr.io/cased/shell:latest
COPY --from=ghcr.io/cased/ssh-oauth-handlers:pr-5 /bin/app /bin/userspace-ssh
COPY --from=ghcr.io/cased/jump:latest /bin/app /bin/jump
ADD entrypoint.sh jump.yml /
ENTRYPOINT [ "/bin/bash", "/entrypoint.sh" ]
CMD []