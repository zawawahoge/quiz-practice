# quiz-practice-proto

protobufでgRPC通信のインターフェイスを定義するリポジトリ。

quiz-practiceのフロント・バックエンド間のインターフェイスをスキーマ言語であるprotobufで定義し、あわよくばgRPCで通信してしまおうという算段である。
backendは、https://github.com/zawawahoge/quiz-practice-api である。

## Why protobuf is used
- 型付きのスキーマ言語を使ってserviceやmessageを定義し、 `protoc` を実行すると、それらを様々な言語で使えるように変換してくれる。


## How to use
```
# at quiz-practice
make proto
```
バックエンドに直接 `pb.go` が生成されるので、conflictを解消してコミットする。
