proto: proto/protobuf/v1/*/*.proto
	docker run --rm -it -v ${PWD}:${PWD} -w ${PWD} znly/protoc \
		-I. \
		--go_out=plugins=grpc:. \
		--grpc-gateway_out=logtostderr=true:. \
		proto/protobuf/v1/service/loginservice.proto \
		proto/protobuf/v1/service/adminservice.proto
	rm -rf api/app/proto
	cp -r proto/protobuf api/app/proto
	rm proto/protobuf/v1/*/*{.pb.go,.pb.gw.go}
	rm api/app/proto/v1/*/*.proto

clean:
	rm -rf ${PWD}/api/{app,reverse-proxy}/proto
