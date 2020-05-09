# API Imageの作成
FROM golang:1.12.4-alpine

# Gopkg.toml 等を事前にコピーして dep ensure が実行できるようにする
COPY ./ /go/go_nuxt

# dep ensure を行うプロジェクトルートに移動する
WORKDIR /go/go_nuxt

# 必要なパッケージをイメージにインストールする
RUN apk update \
  && apk add --no-cache git \
  && export GO111MODULE=on 