default: build

.PHONY: build
build:
	go build -o bin/terraform-provider-spaceliftoutput

.PHONY: test
test:
	go test -v ./...

.PHONY: testacc
testacc:
	TF_ACC=1 go test -v ./...

.PHONY: install
install:
	go install

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: lint
lint: fmt vet

.PHONY: clean
clean:
	rm -rf bin/ 