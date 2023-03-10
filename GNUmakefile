TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=hashicorp.com
NAMESPACE=jonnydgreen
NAME=postman
BINARY=terraform-provider-${NAME}
VERSION=0.2
OS_ARCH=darwin_amd64
GOBIN=$(shell pwd)/bin
CLIENT_PATH=client/postman
PACKAGE_NAME=postman
CLEAN=true
TESTARGS?=""

define local_test_clean
  if [ "$(1)" == "true" ]; then \
		cd examples/testing && \
		rm -rf terraform.tfstate*; \
	fi
endef

.PHONY: default
default: install

.PHONY: build
build:
	go build -o ${BINARY}

.PHONY: release
release:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_darwin_amd64
	GOOS=freebsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_freebsd_386
	GOOS=freebsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_freebsd_amd64
	GOOS=freebsd GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_freebsd_arm
	GOOS=linux GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_linux_386
	GOOS=linux GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_linux_amd64
	GOOS=linux GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_linux_arm
	GOOS=openbsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_openbsd_386
	GOOS=openbsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_openbsd_amd64
	GOOS=solaris GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_solaris_amd64
	GOOS=windows GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_windows_386
	GOOS=windows GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_windows_amd64

.PHONY: install-docs
install-docs:
	@go install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

.PHONY: docs
docs: install-docs
	${GOBIN}/tfplugindocs generate
	${GOBIN}/tfplugindocs generate

.PHONY: gen
gen: 
	go generate ./...

.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

.PHONY: install
install: vendor client gen build docs
	go install .

.PHONY: test
test: 
	go test $(TEST) || exit 1
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

.PHONY: testacc
testacc: 
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

.PHONY: cover
cover: 
	TF_ACC=1 go test $(TEST) -v -coverprofile coverage.out -cover ./postman  -timeout 120m;
	go tool cover -html=coverage.out;

.PHONY: format
format:
	go fmt ./...

.PHONY: local-test-clean
local-test-clean:
	$(call local_test_clean,$(CLEAN))

.PHONY: local-test
local-test: local-test-clean
	cd examples/testing && rm -rf .terraform .terraform.lock.hcl && terraform init && terraform apply --auto-approve

.PHONY: client
client:
	go run cmd/sanitise-openapi-spec/main.go
	rm -rf ${CLIENT_PATH}
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
		-i /local/openapi.yaml \
		-g go \
		-o /local/${CLIENT_PATH} \
		-p=isGoSubmodule=true,packageName=${PACKAGE_NAME} \
		--strict-spec true;
	rm -f ${CLIENT_PATH}/go.*
	rm -rf ${CLIENT_PATH}/api
	rm -rf ${CLIENT_PATH}/test
	rm -rf ${CLIENT_PATH}/docs
	rm -rf ${CLIENT_PATH}/.openapi-generator
	rm -rf ${CLIENT_PATH}/.gitignore
	rm -rf ${CLIENT_PATH}/.travis.yml
	rm -rf ${CLIENT_PATH}/.openapi-generator-ignore
	rm -rf ${CLIENT_PATH}/git_push.sh

.PHONY: vet
vet:
	go vet ./...

.PHONY: pre-commit
pre-commit:
	make
	make vet
	make test
	make gen && git add docs examples
