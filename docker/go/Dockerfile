# goバージョン
FROM golang:1.21.2

# GOPATH(go installでのインストール先のPATH)を通す
ENV GOPATH=/go/src
# WORKSPACEの作成
ENV WORKSPACE=/go/src/app
RUN mkdir -p ${WORKSPACE}

# dockerを開いた時のカレントディレクトリの設定
WORKDIR ${WORKSPACE}

# ファイルをDockerにコピー
COPY . ${WORKSPACE}