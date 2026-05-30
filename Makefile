PROTO_DIR=proto

.PHONY: proto clean-proto

proto:
	@command -v protoc >/dev/null 2>&1 || { echo "protoc not installed"; exit 1; }
	@command -v protoc-gen-go >/dev/null 2>&1 || { echo "protoc-gen-go not installed"; exit 1; }
	@command -v protoc-gen-go-grpc >/dev/null 2>&1 || { echo "protoc-gen-go-grpc not installed"; exit 1; }

	protoc \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(PROTO_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(PROTO_DIR) \
		--go-grpc_opt=paths=source_relative \
		proto/auth/*.proto \
        proto/user/*.proto

clean-proto:
	rm -f $(PROTO_DIR)/*.pb.go
	rm -f $(PROTO_DIR)/*_grpc.pb.go

migrate-up:
	migrate -path db/migrations \
	-database "$(DB_URL)" \
	up

migrate-down:
	migrate -path db/migrations \
	-database "$(DB_URL)" \
	down 1