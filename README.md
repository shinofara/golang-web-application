Nginx Containerと、Go FCGI ContainerをDocker composeを使って立ち上げるサンプル
=========================

# 事前に必要な物

### docker
Macの場合はboot2docker  
Dockerが実行可能なcoreosなどでもOK（https://github.com/coreos/coreos-vagrant）

### docker-compose

https://docs.docker.com/compose/

```
複数のコンテナからなるアプリケーションを
YAMLファイル一つで定義して、簡単なコマンドでアプリケーションの起動や管理ができるツールです。
```

# サンプルを触ってみる

## リポジトリをclone

```
$ git clone https://github.com/shinofara/golang-web-application.git
$ golang-web-application
```

## Containerを立ち上げる

```build an image for your code, and start everything up```

ビルドして、そして全部立ち上げます。

```
$ docker-compose up -d
Recreating demo_go_1...
Recreating demo_nginx_1...
```

## 確認してみましょう

```
$ curl -l http://<docker host serverのIP>:18888
Hello World !
```
