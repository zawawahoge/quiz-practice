# quiz-practice-api

クイズ練習用アプリのバックエンド。GolangでDDDっぽいアーキテクチャで作っているつもり。
まだ何も機能がない。最初はログイン機能から作る予定。

## ディレクトリ構成
```
app
├── cmd
│   └── app
│       └── main.go
├── domain
│   ├── models
│   └── repository
├── handler 
│   └── loginservice
├── infrastructure
│   └── impl
│       ├── loginserviceimpl
│       └── repository
├── proto
│   └── v1
│       └── loginservice
└── usecase
    └── user
```

### `cmd/app`
エントリポイントである `main.go` がある。
軽量なロジックだが、依存性注入などを行う。

### `domain/models`
ユーザ・問題といったドメインモデル

### `domain/repository`
DBやメモリなどとの通信をするためのインターフェイス

### `handler`
リクエストをserviceにハンドルする。gRPCなら消えるけど未実装

### `infrastructure`
各serviceやrepositoryの実装

### `proto/`
serviceのインターフェイスを [protobuf](https://github.com/protocolbuffers/protobuf) で出力したもの。
https://github.com/zawawahoge/quiz-practice-proto のリポジトリで `pb.go` 生成している。
フロントでも `quiz-practice-proto` からインターフェイスを作れるので便利。