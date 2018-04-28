build:
	protoc -I $(GOPATH)/src/github.com/timkellogg/store/user/protos/user \
	$(GOPATH)/src/github.com/timkellogg/store/user/protos/user/user.proto \
	--micro_out=protos/user \
	--go_out=protos/user

clean:
	rm protos/user/user.pb.go protos/user/user.micro.go
