# Makefile for go project
#
# Author: Bernhard Reitinger
#
# Targets:
# 	all: Builds the code
# 	build: Builds the code
# 	fmt: Formats the source files
# 	clean: cleans the code
# 	install: Installs the binaries
# 	test: Runs the tests
#
VERSION := 0.3.2
BUILD := `git rev-parse HEAD`

GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

INSTALL_PATH=/usr/local/bin
MAN_PATH=/usr/local/man

PKGS := $(shell go list ./... | grep -v /vendor)

# Use linker flags to provide version/build settings to the target
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

RXI_SRC = cmd/rxi/*.go
OBJ2REX_SRC = cmd/obj2rex/*.go

TARGETS = rxi obj2rex

all: rxi obj2rex

rxi: $(RXI_SRC)
	$(GOBUILD) -o $@ $(LDFLAGS) $(RXI_SRC)

obj2rex: $(OBJ2REX_SRC)
	$(GOBUILD) -o $@ $(LDFLAGS) $(OBJ2REX_SRC)

clean:
	@rm -f $(TARGETS)

test:
	$(GOTEST) $(PKGS)

install: all
	sudo cp -f $(TARGETS) ${INSTALL_PATH}

uninstall:
	sudo rm -f ${INSTALL_PATH}/rxi
	sudo rm -f ${INSTALL_PATH}/obj2rex

.PHONY: all test install uninstall
