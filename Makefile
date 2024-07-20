
.PHONY: build
build: 
	go build --trimpath -o ./build/  ./cmd/...


.PHONY: clean
clean: 
	rm -rf build

.PHONY: init
init: 
	go install github.com/lyft/protoc-gen-star/protoc-gen-debug@latest

.PHONY: install
install: build  
	cp build/getgopath $(GOPATH)/bin