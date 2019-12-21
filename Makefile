proto: proto/protobuf/v1/*/*.proto
	docker run --rm -it -v ${PWD}:${PWD} -w ${PWD} znly/protoc \
		-I. \
		--go_out=plugins=grpc:. \
		--grpc-gateway_out=logtostderr=true:. \
		proto/protobuf/v1/loginservice/loginservice.proto
	rm -rf ${PWD}/api/app/proto
	cp -r proto/protobuf api/app/proto

clean:
	rm -rf ${PWD}/api/app/proto
	rm proto/protobuf/v1/*/*.pb.go
	rm proto/protobuf/v1/*/*.pb.gw.go
