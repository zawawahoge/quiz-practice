init:
	cp .env.local .env
	GO111MODULE=on go mod download

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
