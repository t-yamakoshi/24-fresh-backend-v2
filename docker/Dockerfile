# 未完成
FROM golang:1.23.1-alpine3.20 as builder

WORKDIR /app

# NOTE: Webp Encode のため以下のパッケージを追加する
# gcc: Cコンパイラ
# musl-dev: 標準Cライブラリ
# libwebp-dev: Webpライブラリ
RUN apk update && apk add git openssh findutils gcc musl-dev libwebp-dev

RUN go env -w GOMODCACHE=/root/.cache/go-build

COPY . ./

RUN --mount=type=cache,target=/root/.cache/go-build go build -o adminapi ./cmd/adminapi
RUN --mount=type=cache,target=/root/.cache/go-build find /root/.cache/go-build/cache/download/ -type f -name *.zip -size +1k | xargs rm -fv
RUN --mount=type=cache,target=/root/.cache/go-build find /root/.cache/go-build/ -type d -name testdata -size +1k | xargs rm -rfv
RUN --mount=type=cache,target=/root/.cache/go-build rm -rfv /root/.cache/go-build/cache/vcs/

FROM alpine:3.20.3

WORKDIR /app

RUN apk add --no-cache ca-certificates libwebp-dev libc6-compat gcompat

COPY --from=builder /app/adminapi .
COPY --from=builder /app/config/ /app/config/

# datadog
COPY --from=gcr.io/datadoghq/serverless-init:1.4.0 /datadog-init /app/datadog-init
ENV DD_OTLP_CONFIG_RECEIVER_PROTOCOLS_GRPC_ENDPOINT="localhost:4317"
ENTRYPOINT ["/app/datadog-init"]
CMD ["/app/adminapi"]

