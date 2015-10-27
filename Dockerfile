FROM gliderlabs/alpine:3.2
MAINTAINER Daisuke Fujita <dtanshi45@gmail.com> (@dtan4)

COPY . /go/src/github.com/dtan4/sintaku
RUN apk update \
      && apk upgrade \
      && apk add git go mercurial \
      && cd /go/src/github.com/dtan4/sintaku \
      && export GOPATH=/go \
      && go get github.com/tools/godep \
      && $GOPATH/bin/godep go build -ldflags "-X main.Version" -o /bin/sintaku \
      && rm -rf /go \
      && apk del --purge go mercurial

EXPOSE 8080

ENTRYPOINT ["/bin/sintaku"]
