# 2020/10/14最新versionを取得
FROM golang:1.19.3-buster
# アップデートとgitのインストール！！
# RUN apt update && apt upgrade
# appディレクトリの作成
RUN mkdir /go/src/app
# ワーキングディレクトリの設定
WORKDIR /go/src/app
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/app

CMD [ "go","run","/main.go" ]