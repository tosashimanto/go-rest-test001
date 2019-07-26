
FROM golang:1.10.2-alpine3.7
LABEL maintainer="tosashimanto"

ADD . ${GOPATH}/src/github.com/tosashimanto/go-rest-test001

RUN go get ./...
# RUN go get github.com/tosashimanto/go-rest-test001




# Gopkg.toml 等を事前にコピーして dep ensure が実行できるようにする
COPY . ${GOPATH}/src/github.com/tosashimanto/go-rest-test001/

# dep ensure を行うプロジェクトルートに移動する
WORKDIR /go/src/github.com/tosashimanto/go-rest-test001

# 必要なパッケージをイメージにインストールする
RUN apk update \
  && apk add --no-cache git \
  && go get -u github.com/golang/dep/cmd/dep \
  && dep ensure


EXPOSE 8080
CMD go-rest-test001
