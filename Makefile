build-proto:
	mkdir -p challenge
	protoc -I/usr/local/include \
		-I./proto \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true:challenge \
		--grpc-gateway_out=logtostderr=true:challenge \
		--go_out=plugins=grpc:./challenge \
		proto/challenge.proto