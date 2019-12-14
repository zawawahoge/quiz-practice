proto: proto/protobuf/*/*.proto
	docker run --rm -v ${PWD}:${PWD} -w ${PWD} znly/protoc \
		-I. \
		--go_out=plugins=grpc:${PWD}/api/app \
		proto/protobuf/loginservice/loginservice.proto
