proto: proto/protobuf/*/*.proto
	docker run --rm -v ${PWD}:${PWD} -w ${PWD} znly/protoc --go_out=plugins=grpc:${PWD}/api/app -I. proto/protobuf/loginservice/loginservice.proto
