FROM golang:1.8

WORKDIR /go/src/github.com/iotv/missouri/
# 2>
RUN apt-get -yq update && \
    apt-get -yq install --no-install-recommends \
        apt-transport-https \
    && \
    echo "deb https://packages.cloud.google.com/apt gcsfuse-jessie main" \
    > /etc/apt/sources.list.d/gcsfuse.list && \
    echo "deb http://ftp.debian.org/debian jessie-backports main" \
    > /etc/apt/sources.list.d/backports.list && \
    curl https://packages.cloud.google.com/apt/doc/apt-key.gpg \
    | apt-key add - && \
    apt-get -yq update && \
    apt-get -yq install --no-install-recommends \
        kmod \
        gcsfuse \
    && \
    apt-get -yqt jessie-backports  install --no-install-recommends \
        ffmpeg \
    && \
    rm -rf /var/lib/apt/lists/*
# </2
# Get glide, dependency manager
RUN go get -v github.com/Masterminds/glide

# Get dependency files and fetch dependencies
COPY ./glide.yaml glide.yaml
COPY ./glide.lock glide.lock
RUN glide install

# Get the rest of the source and build it
COPY ./ .
RUN go build github.com/iotv/missouri

# Setup google cloud storage mounts
ENV GOOGLE_APPLICATION_CREDENTIALS=/go/src/github.com/iotv/missouri/iotv-Raw-Videos-Service-Admin.json
RUN mkdir /data
CMD ./run

# 1 FROM debian

# 3 WORKDIR /opt/
# 4 COPY --from=0 /go/src/github.com/iotv/missouri/missouri .
