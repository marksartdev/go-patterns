.PHONY: default run build lint test cover bench clean full fast

.DEFAULT_GOAL := default

ifdef g
	path := $(g)

	ifdef p
		path := $(path)/$(p)
	else
		path := $(path)/...
	endif
else
	path := ...
endif

default:
ifdef p
	make run
else
	make fast
endif

run:
	make build
	./cmd/$(g)/$(p)/$(p)
	make clean

build:
ifndef g
	@false
endif
ifndef p
	@false
endif

	go build -v -o ./cmd/$(path) ./cmd/$(path)

lint:
	golangci-lint run --enable-all ./cmd/$(path)
	golangci-lint run --enable-all ./pkg/$(path)

test:
	go test -v -cover ./pkg/$(path)

cover:
	go test -v -cover -coverprofile=c.out ./pkg/$(path)
	go tool cover -html=c.out
	rm c.out

bench:
	go test -bench . -benchmem ./pkg/$(path)

clean:
	go clean -r ./cmd/$(path)

full: build lint test clean

fast: lint test
