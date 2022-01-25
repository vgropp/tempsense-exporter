FROM golang:1-alpine as builder
WORKDIR /go/src/github.com/vgropp/tempsense-exporter
RUN apk add build-base linux-headers

COPY . ./
RUN CGO_ENABLED=1 go build -o . ./cmd/tempsense-exporter


FROM alpine:latest

RUN apk add strace
WORKDIR /app
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "$(pwd)" \
    --no-create-home \
    --uid "1000" \
    "appuser"
USER appuser:46
COPY --from=builder /go/src/github.com/vgropp/tempsense-exporter/tempsense-exporter .
ENTRYPOINT ["/app/tempsense-exporter"]
