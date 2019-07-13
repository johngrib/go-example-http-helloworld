FROM golang:1.12

WORKDIR $HOME/go-example-http-helloworld

COPY . .
RUN go mod tidy
CMD ["go", "run", "main.go"]

