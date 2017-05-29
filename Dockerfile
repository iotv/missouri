FROM golang:1.8

WORKDIR /go/src/github.com/iotv/missouri/
RUN go get -v github.com/Masterminds/glide
COPY ./glide.yaml glide.yaml
COPY ./glide.lock glide.lock
RUN glide install
COPY ./ .
RUN go build github.com/iotv/missouri


#//FROM debian

#//RUN apt-get -yq update && \
#//    apt-get -yq install --no-install-recommends \
#//        libav-tools \
#//    && \
#//    rm -rf /var/lib/apt/lists/*
##//WORKDIR /opt/
#//COPY --from=0 /go/src/github.com/iotv/missouri/missouri .
#//CMD ["./missouri"]
