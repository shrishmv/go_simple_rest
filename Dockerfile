FROM golang
RUN mkdir /go/src/go_simple_rest
ADD . /go/src/go_simple_rest/
WORKDIR /go/src/go_simple_rest
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y git
RUN go get github.com/fatih/structs
RUN go get github.com/labstack/echo
RUN go get github.com/labstack/echo/middleware
RUN go get github.com/spf13/viper
RUN go build main.go
CMD ["./main"]