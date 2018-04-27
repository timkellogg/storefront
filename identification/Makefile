build:
	protoc -I $(GOPATH)/src/github.com/timkellogg/store/identification/protos/identification $(GOPATH)/src/github.com/timkellogg/store/identification/protos/identification/identification.proto --go_out=plugins=grpc:protos/identification

clean:
	rm protos/identification/identification.pb.go
