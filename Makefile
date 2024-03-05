generate_grpc_code:
	protoc \
	--go_out=player \
	--go_opt=paths=source_relative \
	--go-grpc_out=player \
	--go-grpc_opt=paths=source_relative \
	player.proto