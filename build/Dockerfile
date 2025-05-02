FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN sh -c 'CURRENT_OS=$(uname -s | tr "[:upper:]" "[:lower:]") && \
    CURRENT_ARCH=$(uname -m) && \
    if [ "$CURRENT_ARCH" = "x86_64" ]; then CURRENT_ARCH="amd64"; fi && \
    if [ "$CURRENT_ARCH" = "aarch64" ]; then CURRENT_ARCH="arm64"; fi && \
    apk add --no-cache git && echo "https://github.com/alibaba/opentelemetry-go-auto-instrumentation/releases/latest/download/otel-$CURRENT_OS-$CURRENT_ARCH" && \
    wget https://github.com/alibaba/opentelemetry-go-auto-instrumentation/releases/latest/download/otel-$CURRENT_OS-$CURRENT_ARCH \
    -O /usr/local/bin/otel \
    && chmod +x /usr/local/bin/otel'

COPY go.mod go.sum ./

RUN go mod download

COPY internal/ internal/
COPY cmd/ cmd/
COPY pkg/ pkg/

ARG COMMITTISH
ARG BUILD_DATE

RUN CGO_ENABLED=0 otel go build -a \
    -ldflags "\
        -s -w \
        -X github.com/adroit-group/go-template/internal/version.committish=${COMMITTISH} \
        -X github.com/adroit-group/go-template/internal/version.buildDate=${BUILD_DATE}" \
    -o bin/service cmd/main.go


FROM cgr.dev/chainguard/static:latest
COPY --from=builder /app/bin/service /usr/local/bin/service
EXPOSE 80
ENTRYPOINT ["service"]
