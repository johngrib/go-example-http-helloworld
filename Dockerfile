FROM golang:1.12

RUN mkdir /hello
COPY main.go /hello
CMD ["go", "run", "/hello/main.go"]
