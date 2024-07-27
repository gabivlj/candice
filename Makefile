DEPS := **/*.go

ifeq ($(OS),Windows_NT)
    MACHINE = WIN32
    ifeq ($(PROCESSOR_ARCHITECTURE),AMD64)
        MACHINE += amd64
    endif
    ifeq ($(PROCESSOR_ARCHITECTURE),x86)
        MACHINE += amd64
    endif
    ifeq ($(PROCESSOR_ARCHITECTURE),ARM64)
        MACHINE += arm64
    endif
    ifeq ($(PROCESSOR_ARCHITECTURE),ARM)
        MACHINE += arm64
    endif
else
    UNAME_S := $(shell uname -s)
    UNAME_P := $(shell uname -p)
    ifeq ($(UNAME_P),x86_64)
        MACHINE += amd64
    endif
    ifneq ($(filter %86,$(UNAME_P)),)
        MACHINE += amd64
    endif
    ifneq ($(filter arm%,$(UNAME_P)),)
        MACHINE += arm64
    endif
endif

PLATFORM ?= linux/$(MACHINE)

.PHONY: build
build: 
	@echo "building candice ${DEPS}"
	cd cmd/build && go build . && bash build.sh 

.PHONY: test
test: build
	@echo "testing candice suite"
	cd internals/tests && go test -v ./.

docker:
	docker build --platform=$(PLATFORM) -t candice:latest .

CMD ?= run . --release
docker-run: docker
	docker run --platform=$(PLATFORM) -v .:/context candice:latest bash -c "cd context && candice $(CMD)"