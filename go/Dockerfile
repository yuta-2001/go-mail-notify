FROM golang:1.21.0 AS builder

COPY . /workspace/go
WORKDIR /workspace/go

ENV ARCH="arm64"

ARG VERSION
RUN go mod download && \
    GOOS=linux GOARCH=${ARCH} CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /workspace/go/main main.go

FROM gcr.io/distroless/static@sha256:ce46866b3a5170db3b49364900fb3168dc0833dfb46c26da5c77f22abb01d8c3 AS prod
COPY --from=builder /workspace/go/main /main
# entrypointはlambdaの設定で上書きされる
ENTRYPOINT ["/main"]

FROM public.ecr.aws/lambda/provided:al2 AS local
COPY --from=builder /workspace/go/main /main
ENTRYPOINT ["/main"]
