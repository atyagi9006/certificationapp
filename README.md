# certificationapp

Prerequisites:

       1. redis dB - installed in Local/ docker image and running on localhost:6379
       2. arango db -installed in local/ docker image and running on localhost:8529 in persistent mode
       3. protoc for compiling  the  grpcproto

Steps to setup: 
1. start arango
    run-arango:
    docker run -e ARANGO_ROOT_PASSWORD=admin -e ARANGO_STORAGE_ENGINE=rocksdb -p 8529:8529 -d  -v  /home/atyagi/goLangWorkSpAce/src/github.com/atyagi9006/certificationapp/db-service/arango_volume:/var/lib/arangodb3 arangodb
2. start redis server
        for local: redis-server --daemonize yes
3. start data-service
        ~/certificationapp> go run data-service/main.go
4. start core-service
        ~/certificationapp> go run core-service/main.go
5. start Portal
        ~/certificationapp/portal>ng serve --open 
6. create db as: test-project-db  and create collections={ user,candidate,question} in arango 
7. Upload Questions to question collection by uploading {maths.json, sciencNature.json,history} one by one only once.
