proto: proto/protobuf/v1/*/*.proto
	docker run --rm -it -v ${PWD}:${PWD} -w ${PWD} znly/protoc \
		-I. \
		--go_out=plugins=grpc:${PWD}/api/app \
		--grpc-gateway_out=logtostderr=true:${PWD} \
		proto/protobuf/v1/loginservice/loginservice.proto
