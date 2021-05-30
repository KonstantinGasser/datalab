proto-common:
	protoc -I. --go_out=$(GOPATH)/src common/common.proto

app-admin:
	protoc -I. --go_out=$(GOPATH)/src service.app.meta.agent/cmd/grpcserver/proto/api.app-meta.proto common/common.proto
	protoc -I. --go-grpc_out=$(GOPATH)/src service.app.meta.agent/cmd/grpcserver/proto/api.app-meta.proto common/common.proto

app-token-issuer:
	protoc -I. --go_out=$(GOPATH)/src service.app.token.agent/cmd/grpcserver/proto/api.app-token.proto common/common.proto
	protoc -I. --go-grpc_out=$(GOPATH)/src service.app.token.agent/cmd/grpcserver/proto/api.app-token.proto common/common.proto

app-configuration:
	protoc -I. --go_out=$(GOPATH)/src service.app.config.agent/cmd/grpcserver/proto/api.app-config.proto common/common.proto
	protoc -I. --go-grpc_out=$(GOPATH)/src service.app.config.agent/cmd/grpcserver/proto/api.app-config.proto common/common.proto

user-admin:
	protoc -I. --go_out=$(GOPATH)/src service.user.meta.agent/cmd/grpcserver/proto/api.user-meta.proto common/common.proto
	protoc -I. --go-grpc_out=$(GOPATH)/src service.user.meta.agent/cmd/grpcserver/proto/api.user-meta.proto common/common.proto

user-auth:
	protoc -I. --go_out=$(GOPATH)/src service.user.auth.agent/cmd/grpcserver/proto/api.user-auth.proto common/common.proto
	protoc -I. --go-grpc_out=$(GOPATH)/src service.user.auth.agent/cmd/grpcserver/proto/api.user-auth.proto common/common.proto

# user-permissions:
# 	protoc -I. --go_out=$(GOPATH)/src service.user-permissions/proto/user-permissions.proto common/common.proto
# 	protoc -I. --go-grpc_out=$(GOPATH)/src service.user-permissions/proto/user-permissions.proto common/common.proto


