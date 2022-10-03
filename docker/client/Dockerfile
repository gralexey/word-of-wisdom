FROM golang:1.19.1-bullseye

COPY . /sources
WORKDIR /sources

RUN go build -o myclient ./cmd/client/...

CMD ["./myclient", "-a", "host.docker.internal:8080", "-c", "20"]
