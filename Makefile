SRC_DIR=adder
DST_DIR=adder

install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go

protoc/grpc:
	protoc -I=${SRC_DIR} --go_out=plugins=grpc:${DST_DIR} ${SRC_DIR}/adder.proto