.PHONY: build clean

clean:
	@rm -rf heap-gen.go cmd/gen/gen

gen: clean
	@rm -rf typed.go
	@go generate generic.go

test: gen
	@go test ./...
