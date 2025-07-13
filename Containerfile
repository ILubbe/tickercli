FROM docker.io/library/golang:1.24 AS builder

WORKDIR /app

COPY go.mod main.go api.key ./
COPY ticker/ ./ticker/
COPY colors/ ./colors/

RUN CGO_ENABLED=0 GOOS=linux go build -o tickercli main.go

FROM docker.io/library/alpine:latest

WORKDIR /app

COPY --from=builder /app/tickercli /app/api.key .

ENV TERM=xterm-256color

ENTRYPOINT ["./tickercli", "-s"]
