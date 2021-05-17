protobuf-common:
	protoc -I. --go_out=$(GOPATH)/src common/common.proto
