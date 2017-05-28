FROM debian

RUN apt-get -yq update && \
    apt-get -yq install --no-install-recommends \
        libav-tools \
    && \
    rm -rf /var/lib/apt/lists/*
