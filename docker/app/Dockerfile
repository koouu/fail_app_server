FROM golang:1.17.6 as builder
WORKDIR /workspace
COPY . /workspace
# alpine でも実行できるように GOOS と CGO_ENABLED を指定
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/api/main.go && chmod +x ./main

FROM alpine:3.15
WORKDIR /app
RUN apk --no-cache add ca-certificates
# root ユーザだとなんでもできてしまうため appuser を作成する
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
COPY --from=builder /workspace/main ./
# コンテナを立ち上げたとき、勝手にWeb サーバーを立ち上げる
CMD ["./main"]