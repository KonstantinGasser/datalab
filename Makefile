
##### Protobuf definitions #####
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

#################################


start-stack:
	echo 'STARTING TO BUILD DOCKER IMAGES...'
	cd ./service.api.bff && ${MAKE} deploy-stack
	cd ./service.app.meta.agent && ${MAKE} deploy-stack
	cd ./service.app.config.agent && ${MAKE} deploy-stack
	cd ./service.app.token.agent && ${MAKE} deploy-stack
	cd ./service.user.meta.agent && ${MAKE} deploy-stack
	cd ./service.user.auth.agent && ${MAKE} deploy-stack
	cd ./service.notification-live && ${MAKE} deploy-stack
	cd ./service.eventmanager.live && ${MAKE} deploy-stack
	echo 'DOCKER IMAGES BUILD!'
	echo '----------------------------------------------'


	
