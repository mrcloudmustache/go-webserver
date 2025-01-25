FROM golang:1.23.5

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-webserver

CMD ["/go-webserver", "-p", "8083"]