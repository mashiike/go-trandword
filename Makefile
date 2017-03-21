all: go-trandword

go-trandword:cmd/trandword/trandword

get-deps:
	glide install

update-deps:
	glide update

test:
	go test -v $$(glide novendor)

SOURCES := $(shell find . -name "*.go" ! -path "./vendor/*" ! -path "./.glide/*")
cmd/trandword/trandword: $(SOURCES) glide.lock
	go build -o cmd/trandword/trandword cmd/trandword/main.go

.PHONY: clean
clean:
	-rm cmd/trandword/trandword
