init:
	cp .env.local .env
	GO111MODULE=on go mod download
	docker-compose exec api_dev go get 'bitbucket.org/liamstask/goose/cmd/goose'

build:
	docker-compose build

run:
	docker-compose up --build

clean:
	docker-compose rm --force --stop

local-test:
	GO111MODULE=on \
	PROJECT_ROOT=${PWD} \
	BUILDER_ROOT=${PWD} \
	go test github.com/zawawahoge/quiz-practice-api/app/...

local-build:
	GO111MODULE=on go build -o bin/echo ./app/cmd/app/*.go

mysql:
	docker-compose exec api_dev mysql -hdb -uroot -ppassword quiz-practice

migrate-up:
	docker-compose exec api_dev goose -path app/db up

migrate-down:
	docker-compose exec api_dev goose -path app/db down

mysql-init:
	docker-compose exec api_dev -t mysql -hdb -uroot -ppassword quiz-practice < app/db/init.sql

look:
	docker-compose exec api_dev bash
