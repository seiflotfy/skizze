GOPATH=$(CURDIR)/vendor:$(CURDIR)

all:
	@GOPATH=$(GOPATH) && \
	  go build -a -v -ldflags '-w' -o ./bin/skizze ./src/skizze

build-dep:
	@go get github.com/constabulary/gb/...

vendor:
	@gb vendor restore

test:
	@GOPATH=$(GOPATH) && go test -race -cover ./src/...

dist: build-dep vendor all

clean:
	@rm ./bin/*

.PHONY: all build-dep vendor test dist clean

