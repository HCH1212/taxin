.PHONY: user-proto
user-proto:
	@protoc --go_out=./api/pb --go-grpc_out=./api/pb api/user.proto

.PHONY: system-proto
system-proto:
	@protoc --go_out=./api/pb --go-grpc_out=./api/pb api/system.proto

.PHONY: pprof-cpu
pprof-cpu:
	@go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30

.PHONY: pprof-mem
pprof-mem:
	@go tool pprof http://localhost:6060/debug/pprof/heap

.PHONY: pprof-goroutine
pprof-goroutine:
	@go tool pprof http://localhost:6060/debug/pprof/goroutine

.PHONY: pprof-block
pprof-block:
	@go tool pprof http://localhost:6060/debug/pprof/block

.PHONY: pprof-mutex
pprof-mutex:
	@go tool pprof http://localhost:6060/debug/pprof/mutex

.PHONY: pprof-trace
pprof-trace:
	@go tool trace http://localhost:6060/debug/pprof/trace

.PHONY: docker-compose-up
docker-compose-up:
	@docker compose up -d

