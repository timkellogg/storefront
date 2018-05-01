build:
	protoc -I $(GOPATH)/src/github.com/timkellogg/store/account/protos/account \
	$(GOPATH)/src/github.com/timkellogg/store/account/protos/account/account.proto \
	--micro_out=protos/account \
	--go_out=protos/account

clean:
	rm protos/account/account.pb.go protos/account/account.micro.go
