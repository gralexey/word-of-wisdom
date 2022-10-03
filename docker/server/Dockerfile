FROM golang:1.19.1-bullseye

COPY . /sources
WORKDIR /sources

RUN go build -o myserver ./cmd/server/...

EXPOSE 8080

CMD ["./myserver"]
