FROM golang:1.4.2

# task runner
RUN go get gopkg.in/godo.v1/cmd/godo

# WAF
RUN go get github.com/zenazn/goji
RUN go get github.com/zenazn/goji/web

WORKDIR /go/src/app

CMD ["/go/bin/godo", "server", "--watch"]

EXPOSE 9000
