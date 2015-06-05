FROM golang:1.4.2

# task runner
RUN go get gopkg.in/godo.v1/cmd/godo

# WAF
RUN go get github.com/gorilla/rpc/v2
RUN go get github.com/gorilla/rpc/v2/json2

WORKDIR /go/src/app

CMD ["/go/bin/godo", "server", "--watch"]

EXPOSE 9000
