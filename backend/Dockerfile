# syntax=docker/dockerfile:1

FROM golang:1.20-buster AS base
RUN set -ex \
    && apt update \
    && apt install -y --no-install-recommends \
        ca-certificates \
        git \
        tzdata \
    && rm -rf /var/lib/apt/lists/*

FROM base AS debugger
WORKDIR /app
COPY . .
RUN go install github.com/go-delve/delve/cmd/dlv@latest
EXPOSE 3001 40000
ENTRYPOINT ["dlv", "debug", "--listen=:40000", "--headless", "--accept-multiclient", "--continue", "--api-version=2"]

FROM base AS development
WORKDIR /app
COPY . .
RUN go install github.com/pilu/fresh@latest
EXPOSE 3001
ENTRYPOINT ["fresh"]

FROM base AS builder
WORKDIR /app
COPY ./backend .
COPY ./.git ./.git
RUN make build

FROM alpine:latest AS productionx
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /app/bin/bosen-backend /bosen-backend
EXPOSE 3001
ENTRYPOINT ["/bosen-backend"]

FROM gcr.io/distroless/base-debian11 as production
WORKDIR /
COPY --from=builder /app/bin/bosen-backend /bosen-backend
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT [ "/bosen-backend" ]