DEPS := **/*.go

.PHONY: build
build: 
	@echo "building candice ${DEPS}"
	cd cmd/build && go build . && bash build.sh 

.PHONY: test
test: build
	@echo "testing candice suite"
	cd internals/tests && go test -v ./.
