run-arango:
	docker run -e ARANGO_ROOT_PASSWORD=admin -e ARANGO_STORAGE_ENGINE=rocksdb -p 8529:8529 -d  -v /home/abhigyan/Workspace/golang_workspace/src/github.com/abhigyandwivedi/tyagi-test-project/db-service/arango_volume:/var/lib/arangodb3 arangodb
proto:
		protoc src/grpcproto/*.proto --go_out=plugins=grpc:.