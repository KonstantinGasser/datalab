server:
	oto -template ./templates/server.go.plush \
	-pkg main \
	-out ./cmd/server/oto.gen.go \
	./pkg/definitions/definitions.go
	gofmt -w ./cmd/server/oto.gen.go ./cmd/server/oto.gen.go	
