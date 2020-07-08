client:
	@echo "--> Generating Go Client files"
	protoc -I protobuf/ --go_out=plugins=grpc:protobuf/ protobuf/kedro.proto
	@echo ""
