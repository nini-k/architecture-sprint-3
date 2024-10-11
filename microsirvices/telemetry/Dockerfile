FROM golang:1.22.5 AS builder
RUN go version

WORKDIR /app
COPY . .
RUN set -Eeux && \
    go mod download && \
    go mod verify

RUN GOOS=linux GOARCH=amd64 \
    go build \
    -o telemetry cmd/telemetry/main.go

FROM scratch
WORKDIR /root/

COPY --from=builder /app/swagger-ui swagger-ui
COPY --from=builder /app/telemetry telemetry

ENTRYPOINT ["./telemetry"]