.PHONY: cputest memtest

cputest:
	@go test -run none -bench . -benchtime 2s -cpuprofile cpu.out

cpupprof: cputest
	@go tool pprof buffer.test cpu.out
	rm -rf buffer.test

memtest:
	@go test -run none -bench . -benchtime 2s -benchmem -memprofile mem.out

mempprof: memtest
	@go tool pprof -alloc_space buffer.test mem.out
	rm -rf buffer.test
