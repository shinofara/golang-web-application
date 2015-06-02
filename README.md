複数のDockerコンテナで造る一つのGo言語WEBアプリケーション開発環境
=========================

# 事前準備

## 必要な物

* [docker](https://docs.docker.com/userguide/)
* [docker-compose](https://docs.docker.com/compose/)
* [docker-machine](https://docs.docker.com/machine/)

## dockerのhost環境をvirtualboxで作成

今回は```docker-machine``` を利用

dockerのhost作成

```
$ docker-machine create -d virtualbox dev
INFO[0000] Creating SSH key...
INFO[0000] Creating VirtualBox VM...
INFO[0009] Starting VirtualBox VM...
INFO[0009] Waiting for VM to start...
INFO[0056] "dev" has been created and is now the active machine. 
```

確認

```
$ docker-machine ls
NAME   ACTIVE   DRIVER       STATE     URL                         SWARM
dev    *        virtualbox   Running   tcp://192.168.99.100:2376
```

dockerコマンドも確認

```
$ docker ps
CONTAINER ID        IMAGE                               COMMAND                CREATED             STATUS              PORTS                         NAMES
```

当然なにもない

## サンプルをクローン

サンプルはこのリポジトリですが

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
Attaching to golangwebapplication_go_1, golangwebapplication_nginx_1 
```

最後にgoをnginxにアタッチと出てます。
これはnginxコンテナにgoコンテナをリンク付けした感じです

## 確認してみましょう

```
$ curl -l http://<docker host serverのIP>/hello/world
Hello World !
```


## 編集してみましょう

vimなどで```main.go```を編集してください。

```
- fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
+ fmt.Fprintf(w, "Hello 2, %s!", c.URLParams["name"])
```

編集したら保存しましょう。閉じちゃってもOKです。

## 確認してみましょう

```
$ curl -l http://<docker host serverのIP>/hello/world
Hello 2 World !
```

## 更新されましたか？

コンテナ内のgoは```go run``` 状態であるはずなのに、なぜ？

それは、goコンテナ内では```go run``` ではなく、```godo``` が実行状態にあるからです。
```godo``` は、gruntのようなものでファイルの変更を検知して、様々なtasksを実行してくれます。

このコンテナではgo run main.goを再実行する様に定義しています。
なので```main.go``` の変更を検知して、内部で```go run main.go``` を再実行してくれたのです。

# 最後に

ここまでで

1. Nginxコンテナ作成
2. Go FCGIコンテナ作成
3. GoのWAFであるgojiの導入
4. ファイルの更新を検知しするタスクランナーgodoの導入

といった４つの事が出来るようになりました。
あとはMysqlコンテナを作って連携すればひと通りのWEBサービス開発を行える状態になりそうです。
