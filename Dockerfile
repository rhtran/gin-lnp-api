# build executable binary
FROM golang:alpine as builder
COPY . $GOPATH/src/gin-lnp-api
WORKDIR $GOPATH/src/gin-lnp-api

RUN mkdir -p /config
#ADD config /config

# git
RUN apk add --no-cache git

## Go dep!
RUN go get github.com/golang/dep/cmd/dep
# Install library dependencies
RUN dep ensure -vendor-only

# build the binary
#RUN go build -o /go/bin/server
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /go/bin/gin-lnp-api
# build a small image
FROM scratch

# copy static executable
COPY --from=builder /go/bin/gin-lnp-api  /go/bin/gin-lnp-api
COPY config/. /config


#ENV PORT 8080
#EXPOSE 8080
ENTRYPOINT ["/go/bin/gin-lnp-api"]
