.PHONY: info
info:
	echo "\n make run : start server on 50051 with db \n make test-service : start service test with server and db \n make test-storage : start storage test with bd"
.PHONY: run
run:
	go run ./cmd/start.go
.PHONY: test-service
test-service:
	go test ./test
.PHONY: test-storage
test-storage:
	go test ./api/storage/library/dbtest
