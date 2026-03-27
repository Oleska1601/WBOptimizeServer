.PHONY: bench-v1 bench-v2 bench-compare 
.PHONY: up down
.PHONY: cpu-v1 cpu-v2 view-cpu-v1 view-cpu-v2 cpu-compare
.PHONY: heap-v1 heap-v2 view-heap-v1 view-heap-v2 heap-compare
.PHONY: allocs-v1 allocs-v2 view-allocs-v1 view-allocs-v2 allocs-compare
.PHONY: trace-v1 trace-v2 view-trace-v1 view-trace-v2

BENCH_DIR = bench_results
BENCH_V1 = $(BENCH_DIR)/v1_bench.txt
BENCH_V2 = $(BENCH_DIR)/v2_bench.txt

load:
	mkdir -p profiles bench_results
	@go run cmd/load/main.go -url http://localhost:8081/api/v1/cpu -c 5 -d 30 &
	@go run cmd/load/main.go -url http://localhost:8081/api/v1/memory -c 5 -d 30 &
	@go run cmd/load/main.go -url http://localhost:8082/api/v2/cpu -c 5 -d 30 &
	@go run cmd/load/main.go -url http://localhost:8082/api/v2/memory -c 5 -d 30 &

up:
	docker compose -f docker-compose.yaml --env-file .env up -d --build

down:
	docker compose -f docker-compose.yaml down

bench-v1:
	go test -bench=. -benchmem ./internal/service/v1/...

bench-v2:
	go test -bench=. -benchmem ./internal/service/v2/... 

bench-all: bench-v1 bench-v2

bench-save-v1: 
	go test -bench=. -benchmem -count=5 ./internal/service/v1/... > bench_results/v1_bench.txt

bench-save-v2: 
	go test -bench=. -benchmem -count=5 ./internal/service/v2/... > bench_results/v2_bench.txt

bench-save-all: bench-save-v1 bench-save-v2

benchstat:
	@echo "Benchmark comparison v1 vs v2"
	@benchstat $(BENCH_V1) $(BENCH_V2)

cpu-v1: 
	curl -o profiles/v1_cpu.prof "http://localhost:8081/api/v1/debug/pprof/profile?seconds=30"

cpu-v2: 
	curl -o profiles/v2_cpu.prof "http://localhost:8082/api/v2/debug/pprof/profile?seconds=30"

view-cpu-v1:
	go tool pprof -http=:8083 profiles/v1_cpu.prof

view-cpu-v2:
	go tool pprof -http=:8084 profiles/v2_cpu.prof

view-cpu-compare:
	go tool pprof -http=:8085 -diff_base=profiles/v1_cpu.prof profiles/v2_cpu.prof

heap-v1: 
	curl -o profiles/v1_heap.prof "http://localhost:8081/api/v1/debug/pprof/heap"

heap-v2: 
	curl -o profiles/v2_heap.prof "http://localhost:8082/api/v2/debug/pprof/heap"

view-heap-v1:
	go tool pprof -http=:8086 profiles/v1_heap.prof

view-heap-v2:
	go tool pprof -http=:8087 profiles/v2_heap.prof

view-heap-compare:
	go tool pprof -http=:8088 -diff_base=profiles/v1_heap.prof profiles/v2_heap.prof

allocs-v1: 
	curl -o profiles/v1_allocs.prof "http://localhost:8081/api/v1/debug/pprof/allocs"

allocs-v2: 
	curl -o profiles/v2_allocs.prof "http://localhost:8082/api/v2/debug/pprof/allocs"

view-allocs-v1:
	go tool pprof -http=:8089 profiles/v1_allocs.prof

view-allocs-v2:
	go tool pprof -http=:8090 profiles/v2_allocs.prof

trace-v1: 
	curl -o profiles/v1_trace.out "http://localhost:8081/api/v1/debug/pprof/trace?seconds=20"

trace-v2: 
	curl -o profiles/v2_trace.out "http://localhost:8082/api/v2/debug/pprof/trace?seconds=20"

view-trace-v1:
	go tool trace profiles/v1_trace.out

view-trace-v2:
	go tool trace profiles/v2_trace.out

run: load cpu-v1 cpu-v2 heap-v1 heap-v2 allocs-v1 allocs-v2 trace-v1 trace-v2

