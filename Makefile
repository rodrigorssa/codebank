gen:
	protoc --proto_path=infra/grpc infra/grpc/proto/*.proto --go_out=infra/ --go-grpc_out=infra/