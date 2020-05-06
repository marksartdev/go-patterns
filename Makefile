.PHONY: run build lint test clean all

.DEFAULT_GOAL := run

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

run:
	make build
	./cmd/$(g)/$(p)/$(p)

build:
ifndef g
	@false
endif
ifndef p
	@false
endif

	go build -v -o ./cmd/$(g)/$(p) ./cmd/$(g)/$(p)

lint:
	golangci-lint run --enable-all ./cmd/$(path)
	golangci-lint run --enable-all ./pkg/$(path)

test:
	go test -v -cover ./pkg/$(path)

clean:
	go clean -r ./cmd/$(path)

all:
ifdef g
ifdef p
	make build
endif
endif

	make lint
	make test
	make clean
