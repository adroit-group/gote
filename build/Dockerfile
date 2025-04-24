FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY internal/ internal/
COPY cmd/ cmd/
COPY pkg/ pkg/

ARG COMMITTISH
ARG BUILD_DATE

RUN CGO_ENABLED=0 go build -a -ldflags "\
		-s -w \
		-X github.com/adroit-group/go-template/pkg/version.Committish=${COMMITTISH} \
		-X github.com/adroit-group/go-template/pkg/version.BuildDate=${BUILD_DATE}" \
		-o bin/service cmd/main.go


FROM cgr.dev/chainguard/static:latest
COPY --from=builder /app/bin/service /usr/local/bin/service
EXPOSE 80
ENTRYPOINT ["service"]
