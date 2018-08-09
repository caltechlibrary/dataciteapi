#
# Simple Makefile for conviently testing, building and deploying experiment.
#
PROJECT = dataciteapi

VERSION = $(shell grep -m 1 'Version =' $(PROJECT).go | cut -d\`  -f 2)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

PKGASSETS = $(shell which pkgassets)

PROJECT_LIST = dataciteapi

OS = $(shell uname)

EXT = 
ifeq ($(OS), Windows)
	EXT = .exe
endif


dataciteapi$(EXT): bin/dataciteapi$(EXT)

cmd/dataciteapi/assets.go:
	pkgassets -o cmd/dataciteapi/assets.go -p main -ext=".md" -strip-prefix="/" -strip-suffix=".md" Examples how-to Help docs/dataciteapi
	git add cmd/dataciteapi/assets.go

bin/dataciteapi$(EXT): dataciteapi.go cmd/dataciteapi/dataciteapi.go cmd/dataciteapi/assets.go
	go build -o bin/dataciteapi$(EXT) cmd/dataciteapi/dataciteapi.go cmd/dataciteapi/assets.go

build: $(PROJECT_LIST)

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
	cd dist && zip -r $(PROJECT)-$(VERSION)-linux-amd64.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

dist/windows-amd64:
	mkdir -p dist/bin
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/dataciteapi.exe cmd/dataciteapi/dataciteapi.go cmd/dataciteapi/assets.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-windows-amd64.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

dist/macosx-amd64:
	mkdir -p dist/bin
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/dataciteapi cmd/dataciteapi/dataciteapi.go cmd/dataciteapi/assets.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-macosx-amd64.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

dist/raspbian-arm7:
	mkdir -p dist/bin
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/dataciteapi cmd/dataciteapi/dataciteapi.go cmd/dataciteapi/assets.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-raspbian-arm7.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

distribute_docs:
	if [ -d dist ]; then rm -fR dist; fi
	mkdir -p dist
	cp -v README.md dist/
	cp -v LICENSE dist/
	cp -v INSTALL.md dist/
	bash package-versions.bash > dist/package-versions.txt

update_version:
	./update_version.py --yes

release: clean dataciteapi.go distribute_docs dist/linux-amd64 dist/windows-amd64 dist/macosx-amd64 dist/raspbian-arm7

status:
	git status

save:
	if [ "$(msg)" != "" ]; then git commit -am "$(msg)"; else git commit -am "Quick Save"; fi
	git push origin $(BRANCH)

publish:
	bash mk-website.bash
	bash publish.bash

