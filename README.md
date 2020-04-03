# Go API

## 起動
```
$ sudo docker build -t go-api .
$ sudo docker run --rm -p 8080:80 -it go-api:latest
```

## 確認
```
$ curl localhost:8080
{"message":"Hello World!"}

$ curl localhost:8080/plus/1
{"number":2}

$ curl localhost:8080/get?name=jack
{"message":"Hello, jack!"}

$ curl -X POST -d "name=jack" localhost:8080/post
{"message":"Hello, jack!"}
```

## 参考
+ [はじめてのGo](http://gihyo.jp/dev/feature/01/go_4beginners)
+ [project-layout](https://github.com/golang-standards/project-layout)
+ [gorilla/mux](https://github.com/gorilla/mux)
