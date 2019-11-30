FROM golang:1.13 as builder

ENV GO111MODULE on
ENV PROJECT_ROOT /go/src/github.com/zawawahoge/quiz-practice/api
WORKDIR $PROJECT_ROOT

RUN apt-get update -y && apt-get install default-mysql-client -y

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o build ./app/cmd/app

FROM alpine:3.10.3 as release
ENV BUILDER_ROOT /go/src/github.com/zawawahoge/quiz-practice
ENV PROJECT_ROOT /
COPY --from=builder $BUILDER_ROOT/build $PROJECT_ROOT/app
COPY --from=builder $BUILDER_ROOT/.env $PROJECT_ROOT/.env
CMD ["/app"]