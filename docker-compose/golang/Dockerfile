FROM golang
LABEL maintainer="tosashimanto"

ADD . ${GOPATH}/src/github.com/FLECT-DEV-TEAM/go-rest-test001
# RUN go get ./...
# RUN go get github.com/FLECT-DEV-TEAM/go-rest-test001
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init
RUN dep ensure

EXPOSE 8080
CMD go-rest-test001
