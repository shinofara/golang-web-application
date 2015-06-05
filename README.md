NginxとGoコンテナでjsonrpc apiサーバを開発環境を造る
=========================

# 履歴

* [NginxとGoコンテナでWEBアプリケーションをつくる](https://github.com/shinofara/golang-web-application)
* [フレームワークでGojiを採用、Godoで監視ビルド対応](https://github.com/shinofara/golang-web-application/tree/goji)
* そして今回

# 事前準備

## 必要な物

* [docker](https://docs.docker.com/userguide/)
* [docker-compose](https://docs.docker.com/compose/)
* [docker-machine](https://docs.docker.com/machine/)

今回はdockerは動いてる前提で進めていきます。


## 登場するGoパッケージ

jsonrpcは標準ライブラリではなく、Gorillaと呼ばれるWAFを使ってみます。

* go get github.com/gorilla/rpc/v2       
* go get github.com/gorilla/rpc/v2/json2 
* go get gopkg.in/godo.v1/cmd/godo


## サンプルをクローン

```
$ git clone https://github.com/shinofara/golang-web-application.git
$ golang-web-application
$ git checkout gorira-jsonrpc-2
```

## Containerを立ち上げる

```
$ docker-compose up -d
```

最後にgoをnginxにアタッチと出てます。
これはnginxコンテナにgoコンテナをリンク付けした感じです

## 確認してみましょう

```
$ curl -H "Content-Type: application/json" -d '{"jsonrpc":"2.0", "id":"test", "method": "Counter.Get", "params": {}}' -s http://192.168.99.100/jsonrpc/ 
{"jsonrpc":"2.0","result":{"Count":0},"id":"test"} 
```

jsonrpc 2.0protocolでやりとりが出来ました。

## APIのmethodの追加

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

# 今回の大きな変更点

今回はGorilla,http,fcgiの連携だったりNginxの設定変更を行いました。