generate_proto_files:
	protoc \
	--go_out=greeting \
	--go_opt=paths=source_relative \
	--go-grpc_out=greeting \
	--go-grpc_opt=paths=source_relative \
	proto/greeting_service.proto