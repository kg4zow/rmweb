########################################
# Filename of the final executable

NAME    := rmweb

########################################
# Pre-calculate values used in 'version' output so that if building multiple
# arches, they all have the same version/time/etc.

NOW     := $(shell date -u +%Y-%m-%dT%H:%M:%SZ )
VERSION := $(shell head -1 version.txt )
HASH    := $(shell git rev-parse --short HEAD 2>/dev/null || true )
DESC    := $(shell git describe --dirty --broken --long 2>/dev/null || true )

########################################
# Other automatic variables

SOURCES := $(shell find . -name "*.go" )
MYOS    := $(shell go env GOOS )
MYARCH  := $(shell go env GOARCH )

########################################
# Which OS/ARCH combinations will be built by 'make all'
# - run 'go tool dist list' to see all available combinations

ALL_ARCHES := darwin/amd64 darwin/arm64 \
		linux/386 linux/amd64 \
		windows/386 windows/amd64

###############################################################################
#
# First/default target: build for *this* machine's OS/ARCH

$(NAME): out/$(NAME)-$(MYOS)-$(MYARCH)
	ln -sf "out/$(NAME)-$(MYOS)-$(MYARCH)" "$(NAME)"

########################################
# Build OS-ARCH/NAME for all combinations listed in ARCHES above
# - if OS is "windows", add ".exe" to the name

all: $(foreach A,$(ALL_ARCHES),out/$(NAME)-$(subst /,-,$(A))$(if $(A:windows%=),,.exe))
	ln -sf "out/$(NAME)-$(MYOS)-$(MYARCH)" "$(NAME)"

########################################
# Remove all previously compiled binaries and symlinks

clean:
	rm -rf $(NAME) out

###############################################################################
#
# How to build OS-ARCH/NAME for any OS/ARCH combination

out/$(NAME)-%: go.mod Makefile version.txt $(SOURCES)
	GOOS=$(shell echo "$@" | awk -F '[\-/]' '{print $$3}' ) \
	GOARCH=$(shell echo "$@" | awk -F '[\-\./]' '{print $$4}' ) \
	go build -o $@ \
	-ldflags "-X main.prog_name=$(NAME) \
	        -X main.prog_version=$(VERSION) \
	        -X main.prog_date=$(NOW) \
	        -X main.prog_hash=$(HASH) \
	        -X main.prog_desc=$(DESC)"

########################################
# Specific rule for reMarkable 2 tablet ... linux/arm with GOARM=7
# - add "linux/arm" to ALL_ARCHES to build this

out/$(NAME)-linux-arm: go.mod Makefile version.txt $(SOURCES)
	GOOS=linux GOARCH=arm GOARM=7 \
	go build -o $@ \
	-ldflags "-X main.prog_name=$(NAME) \
	        -X main.prog_version=$(VERSION) \
	        -X main.prog_date=$(NOW) \
	        -X main.prog_hash=$(HASH) \
	        -X main.prog_desc=$(DESC)"
