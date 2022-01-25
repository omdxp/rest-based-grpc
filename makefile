create:
	protoc --proto_path=proto proto/*.proto --go_out=gen/proto/
	protoc --proto_path=proto proto/*.proto --go-grpc_out=gen/proto/
	protoc -I . --grpc-gateway_out ./gen \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		proto/test.proto

clean:
	rm -rf gen/proto/

google_annotations:
	mkdir -p google/api
	curl -L https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto > google/api/annotations.proto
	curl -L https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto > google/api/http.proto

run_server:
	go run server/server.go

run_client:
	go run client/client.go