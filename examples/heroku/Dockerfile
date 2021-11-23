FROM golang:1.16 as builder

# Install upx (upx.github.io) to compress the compiled action
RUN apt-get update && apt-get -y install upx

# Turn on Go modules support and disable CGO
ENV GO111MODULE=on CGO_ENABLED=0

# Copy all the files from the host into the container
WORKDIR /src
COPY . .

# Compile the action - the added flags instruct Go to produce a
# standalone binary
RUN find . &&  \
  go build \
  -a \
  -trimpath \
  -ldflags "-s -w -extldflags '-static'" \
  -installsuffix cgo \
  -tags netgo \
  -o /bin/app \
  .

# Strip any symbols - this is not a library
RUN strip /bin/app

# Compress the compiled action
RUN upx -q -9 /bin/app

FROM ghcr.io/cased/shell:pr-1520
ENTRYPOINT [ "/bin/bash", "/entrypoint.sh" ]
CMD []
RUN curl https://cli-assets.heroku.com/install.sh | sh
ADD entrypoint.sh prompts.json heroku-bashrc /
COPY --from=builder /bin/app /bin/heroku-ssh
ENV CASED_SHELL_HOST_FILE=/prompts.json