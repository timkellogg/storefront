build:
	protoc -I $(GOPATH)/src/github.com/timkellogg/store/identification/ $(GOPATH)/src/github.com/timkellogg/store/identification/identification.proto --go_out=plugins=grpc:.

clean:
	rm identification.pb.go

