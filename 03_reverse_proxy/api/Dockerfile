#-------------------------------------------------------
# 開発環境
#-------------------------------------------------------
FROM golang:1.21.5-alpine3.19 as builder

# 依存パッケージのインストール
RUN apk update

# Gin APIサーバーのビルド
COPY gin /gin
WORKDIR /gin
RUN go build main.go


#-------------------------------------------------------
# 製品環境
#-------------------------------------------------------
FROM gcr.io/distroless/static-debian12:nonroot as releaser

# Gin APIサーバーのインストール
COPY --from=builder /gin /gin
CMD ["/gin/main"]
