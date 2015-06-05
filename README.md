NginxとGoコンテナでjsonrpc apiサーバを開発環境を造る
=========================

# 事前準備

## 必要な物

* [docker](https://docs.docker.com/userguide/)
* [docker-compose](https://docs.docker.com/compose/)
* [docker-machine](https://docs.docker.com/machine/)

## 登場するGoパッケージ

jsonrpcは標準ライブラリではなく、Gorillaと呼ばれるWAFを使ってみます。

* go get github.com/gorilla/rpc/v2       
* go get github.com/gorilla/rpc/v2/json2 
* go get gopkg.in/godo.v1/cmd/godo


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
$ git checkout gorira-jsonrpc-2
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
$ curl -H "Content-Type: application/json" -d '{"jsonrpc":"2.0", "id":"test", "method": "Counter.Get", "params": {}}' -s http://192.168.99.100/jsonrpc/ 
{"jsonrpc":"2.0","result":{"Count":0},"id":"test"} 
```

jsonrpc 2.0protocolでやりとりが出来ました。

## methodの追加

```
s := rpc.NewServer()                                                                  
s.RegisterCodec(json2.NewCustomCodec(&rpc.CompressionSelector{}), "application/json") 
s.RegisterService(new(Counter), "")                                                   
http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./"))))              
http.Handle("/jsonrpc/", s)                                                           
listener, _ := net.Listen("tcp", ":9000")                                             
                                                                                      
log.Fatal(fcgi.Serve(listener, nil))                                                  
```

コードを見てもらえればと思いますが、```s.RegisterService``` でmethodの登録を行っています。