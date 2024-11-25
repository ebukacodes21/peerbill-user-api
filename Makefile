proto:
	rm -f pb/*.go
	rm -f doc/swagger/*swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=true \
	proto/*.proto
	statik -src=./doc/swagger -dest=./doc

evans:
	docker run --rm -it -v "/Users/george/workspace/peerbill-user-api:/mount:ro" \
    ghcr.io/ktr0731/evans:latest \
    --path /mount/proto/ \
    --proto peerbill_user.proto \
    --host host.docker.internal \
    --port 9092 \
    repl

server:
	go run main.go

.PHONY: proto server evans