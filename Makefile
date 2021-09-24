#
# Simple Makefile for conviently testing, building and deploying experiment.
#
PROJECT = dataciteapi

VERSION = $(shell grep '"version":' codemeta.json | cut -d \" -f 4)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

PKGASSETS = $(shell which pkgassets)

PROJECT_LIST = dataciteapi

OS = $(shell uname)

EXT = 
ifeq ($(OS), Windows)
	EXT = .exe
endif

build: version.go $(PROJECT_LIST)

dataciteapi$(EXT): bin/dataciteapi$(EXT)

cmd/dataciteapi/assets.go:
	pkgassets -o cmd/dataciteapi/assets.go -p main -ext=".md" -strip-prefix="/" -strip-suffix=".md" Examples how-to Help docs/dataciteapi
	git add cmd/dataciteapi/assets.go

bin/dataciteapi$(EXT): dataciteapi.go cmd/dataciteapi/dataciteapi.go cmd/dataciteapi/assets.go
	go build -o bin/dataciteapi$(EXT) cmd/dataciteapi/dataciteapi.go cmd/dataciteapi/assets.go


version.go: .FORCE
	@echo "package $(PROJECT)" >version.go
	@echo '' >>version.go
	@echo 'const Version = "v$(VERSION)"' >>version.go
	@echo '' >>version.go
	@git add version.go


install: 
	env GOBIN=$(GOPATH)/bin go install cmd/dataciteapi/dataciteapi.go cmd/dataciteapi/assets.go

website: page.tmpl README.md nav.md INSTALL.md LICENSE css/site.css
	bash mk-website.bash

test: clean bin/dataciteapi$(EXT)
	go test -mailto="jane.doe@example.edu"

format:
	gofmt -w dataciteapi.go
	gofmt -w dataciteapi_test.go
	gofmt -w cmd/dataciteapi/dataciteapi.go

lint:
	golint dataciteapi.go
	golint dataciteapi_test.go
	golint cmd/dataciteapi/dataciteapi.go

clean: 
	if [ "$(PKGASSETS)" != "" ]; then bash rebuild-assets.bash; fi
	if [ -f index.html ]; then rm *.html; fi
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi
	if [ -d man ]; then rm -fR man; fi
	if [ -d testdata ]; then rm -fR testdata; fi

man: build
	mkdir -p man/man1
	bin/dataciteapi -generate-manpage | nroff -Tutf8 -man > man/man1/dataciteapi.1

dist/linux-amd64:
	mkdir -p dist/bin
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/dataciteapi cmd/dataciteapi/dataciteapi.go cmd/dataciteapi/assets.go
	cd dist && zip -r $(PROJECT)-v$(VERSION)-linux-amd64.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

dist/windows-amd64:
	mkdir -p dist/bin
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/dataciteapi.exe cmd/dataciteapi/dataciteapi.go cmd/dataciteapi/assets.go
	cd dist && zip -r $(PROJECT)-v$(VERSION)-windows-amd64.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

dist/macos-amd64:
	mkdir -p dist/bin
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/dataciteapi cmd/dataciteapi/dataciteapi.go cmd/dataciteapi/assets.go
	cd dist && zip -r $(PROJECT)-v$(VERSION)-macos-amd64.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

dist/macos-arm64:
	mkdir -p dist/bin
	env  GOOS=darwin GOARCH=arm64 go build -o dist/bin/dataciteapi cmd/dataciteapi/dataciteapi.go cmd/dataciteapi/assets.go
	cd dist && zip -r $(PROJECT)-v$(VERSION)-macos-arm64.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

dist/raspbian-arm7:
	mkdir -p dist/bin
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/dataciteapi cmd/dataciteapi/dataciteapi.go cmd/dataciteapi/assets.go
	cd dist && zip -r $(PROJECT)-v$(VERSION)-raspbian-arm7.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

distribute_docs:
	if [ -d dist ]; then rm -fR dist; fi
	mkdir -p dist
	cp -v README.md dist/
	cp -v LICENSE dist/
	cp -v INSTALL.md dist/

release: clean dataciteapi.go distribute_docs dist/linux-amd64 dist/windows-amd64 dist/macos-amd64 dist/macos-arm64 dist/raspbian-arm7

status:
	git status

save:
	if [ "$(msg)" != "" ]; then git commit -am "$(msg)"; else git commit -am "Quick Save"; fi
	git push origin $(BRANCH)

publish:
	bash mk-website.bash
	bash publish.bash


.FORCE:
