proto-common:
	protoc -I. --go_out=$(GOPATH)/src common/common.proto

app-admin:
	protoc -I. --go_out=$(GOPATH)/src service.app-administer/proto/app-administer.proto common/common.proto
	protoc -I. --go-grpc_out=$(GOPATH)/src service.app-administer/proto/app-administer.proto common/common.proto

app-token-issuer:
	protoc -I. --go_out=$(GOPATH)/src service.app-token-issuer/proto/app-token-issuer.proto
	protoc -I. --go-grpc_out=$(GOPATH)/src service.app-token-issuer/proto/app-token-issuer.proto

app-configuration:
	protoc -I. --go_out=$(GOPATH)/src service.app-configuration/proto/app-configuration.proto common/common.proto
	protoc -I. --go-grpc_out=$(GOPATH)/src service.app-configuration/proto/app-configuration.proto common/common.proto

user-admin:
	protoc -I. --go_out=$(GOPATH)/src service.user-administer/proto/user-administer.proto common/common.proto
	protoc -I. --go-grpc_out=$(GOPATH)/src service.user-administer/proto/user-administer.proto common/common.proto

user-auth:
	protoc -I. --go_out=$(GOPATH)/src service.user-authentication/proto/user-authentication.proto common/common.proto
	protoc -I. --go-grpc_out=$(GOPATH)/src service.user-authentication/proto/user-authentication.proto common/common.proto

user-permissions:
	protoc -I. --go_out=$(GOPATH)/src service.user-permissions/proto/user-permissions.proto common/common.proto
	protoc -I. --go-grpc_out=$(GOPATH)/src service.user-permissions/proto/user-permissions.proto common/common.proto