VERSION = 1.1.1

PACKAGE_NAME := git-glimpse
BIN_DIR := bin
GO_FILES := $(shell find . -type f -name '*.go')

BUILD_FLAGS := -ldflags "-X main.ExecutableName=${PACKAGE_NAME} -X main.Version=${VERSION} -s -w" -tags static

.PHONY: all
all: ${BIN_DIR}/${PACKAGE_NAME}

.PHONY: clean
clean:
	rm -f ${BIN_DIR}/${PACKAGE_NAME}

# Build the executable
${BIN_DIR}/${PACKAGE_NAME}: ${GO_FILES}
	go build ${BUILD_FLAGS} -o $@

.PHONY: install
install:
	go install ${BUILD_FLAGS}

# TODO investigate on uninstall
